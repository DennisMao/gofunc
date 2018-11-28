// Elasticseach scroll 分页查询测试
package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"os"
	"sync"
	"time"

	es "gopkg.in/olivere/elastic.v5"
)

// 参数配置
const (
	EsUrl   = "XXX"
	EsIndex = "testdb-20181101"
	EsType  = "test_type"
)

// 数据模型
type EsTestData struct {
	Id         string `json:"id"`
	EsType     string `json:"es_type"`
	EsIndex    string `json:"es_index"`
	CreateTime int64  `json:"create_time"`
	Content    string `json:"types"`
	Uid        int64  `json:"uid"`
}

// 样例1 FromSize
// 获取指定创建时间点下所有用户数据
// @createTime 创建时间(用于指定搜索点)
// @from 起始数据点
// @size 最大偏移量
func FromSize(createTime, from, size int64) ([]EsTestData, error) {
	retry := es.NewBackoffRetrier(es.ZeroBackoff{})
	client, err := es.NewClient(
		es.SetURL(EsUrl),
		es.SetSniff(false),
		es.SetTraceLog(log.New(os.Stdout, "ES ", log.LstdFlags)),
		es.SetRetrier(retry),
	)
	if err != nil {
		panic(err)
	}

	// 查询条件
	qTx := es.NewBoolQuery()
	qTx.Must(es.NewTermQuery("create_time", createTime))

	// 协程控制
	ctx, _ := context.WithCancel(context.Background())
	tStart := time.Now()

	results, err := client.Search().Index(EsIndex).Type(EsType).Query(qTx).From(int(from)).Size(int(size)).Do(ctx)
	if err == io.EOF {
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	var ret []EsTestData

	for i, _ := range results.Hits.Hits {
		var d EsTestData
		if err := json.Unmarshal(*(results.Hits.Hits[i].Source), &d); err != nil {
			continue
		}

		ret = append(ret, d)
	}

	log.Printf("[Mget] 获取完成,共计获取到数据数量:%d 耗时:%s \n", results.Hits.TotalHits, time.Now().Sub(tStart).String())
	return ret, nil
}

// 样例2 SearchAfter
// 获取指定创建时间点下所有用户数据
// @createTime 创建时间
func SearchAfter(createTime int64) ([]EsTestData, error) {
	retry := es.NewBackoffRetrier(es.ZeroBackoff{})
	client, err := es.NewClient(
		es.SetURL(EsUrl),
		es.SetSniff(false),
		es.SetTraceLog(log.New(os.Stdout, "ES ", log.LstdFlags)),
		es.SetRetrier(retry),
	)
	if err != nil {
		panic(err)
	}

	qTx := es.NewBoolQuery()
	qTx.Must(es.NewTermQuery("create_time", createTime))

	//协程控制
	ctx, _ := context.WithCancel(context.Background())
	tStart := time.Now()

	var hits []*json.RawMessage
	lastUid := int64(0)
	cnt := int64(0)

	for {
		searchAfter := es.NewSearchSource().Query(qTx).SearchAfter(lastUid).Sort("uid", true)
		results, err := client.Search().Index(EsIndex).Type(EsType).SearchSource(searchAfter).Size(5000).Do(ctx)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("[SearchAfter] 查询失败,错误原因:", err.Error())
			return nil, err
		}

		if len(results.Hits.Hits) == 0 {
			break
		}

		//解析最后一个
		var lastCell EsTestData
		err = json.Unmarshal(*(results.Hits.Hits[len(results.Hits.Hits)-1]).Source, &lastCell)
		if err != nil {
			log.Println("[SearchAfter] 最后一个解析失败")
			break
		}

		lastUid = lastCell.Uid
		cnt += results.Hits.TotalHits

		//收集结果
		for i, _ := range results.Hits.Hits {
			hits = append(hits, results.Hits.Hits[i].Source)
		}
	}

	var ret []EsTestData
	for i, _ := range hits {
		var d EsTestData
		if err := json.Unmarshal(*(hits[i]), &d); err != nil {
			continue
		}

		ret = append(ret, d)

	}

	log.Println("[SearchAfter] 获取完成,共计获取到数据数量:%d 有效数据数据量:%d 耗时:%s", cnt, len(ret), time.Now().Sub(tStart).String())
	return ret, nil
}

// 样例3 Scroll
// 获取指定创建时间点下所有用户数据
// @createTime 创建时间
func Scroll(createTime int64) ([]EsTestData, error) {
	retry := es.NewBackoffRetrier(es.ZeroBackoff{})
	client, err := es.NewClient(
		es.SetURL(EsUrl),
		es.SetSniff(false),
		es.SetTraceLog(log.New(os.Stdout, "ES ", log.LstdFlags)),
		es.SetRetrier(retry),
	)
	if err != nil {
		panic(err)
	}

	//筛选条件 refer =="tx"
	qTx := es.NewBoolQuery()
	qTx.Must(es.NewTermQuery("create_time", createTime))

	//解析通道
	//解析速度需要快于读取速度
	hitChan := make(chan json.RawMessage, 10)

	//协程控制
	ctx, _ := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	tStart := time.Now()

	//Scroll读取协程
	wg.Add(1)
	go func() error {
		defer close(hitChan)
		defer wg.Done()

		scroll := client.Scroll(EsIndex).Type(EsType).Query(qTx).Size(200)
		for {
			results, err := scroll.Do(ctx)
			if err == io.EOF {
				return nil
			}
			if err != nil {
				return err
			}

			// 通过channel发送数据到解析协程
			for _, hit := range results.Hits.Hits {
				select {
				case hitChan <- *hit.Source:
				case <-ctx.Done():
					return ctx.Err()
				}
			}
		}
		return nil
	}()

	var ret []EsTestData
	retRw := sync.RWMutex{}

	//数据处理协程  5个
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() error {
			defer wg.Done()
			for hit := range hitChan {
				var d EsTestData
				err := json.Unmarshal(hit, &d)
				if err != nil {
					log.Println("[Scroll] [解析] 错误数据:", string(hit))
					continue
				}

				retRw.Lock()
				ret = append(ret, d)
				retRw.Unlock()
			}

			return nil
		}()
	}

	wg.Wait()

	log.Printf("[Scroll] 获取完成,共计获取到数据数量:%d 耗时:%s \n", len(ret), time.Now().Sub(tStart).String())
	return ret, nil
}

func main() {
	FromSize(1541052000, 0, 10000)
	SearchAfter(1541052000)
	Scroll(1541052000)
}
