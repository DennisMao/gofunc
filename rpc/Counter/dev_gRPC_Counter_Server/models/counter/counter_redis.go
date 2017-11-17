/*
计数器-Redis层
*/
package counter

import (
	"errors"
	"fmt"
	"strconv"
	"sync"
)

const (
	REDIS_PORT           = "6379"
	REDIS_COLLECTIONNAME = "counter"
	REDIS_PASSWORLD      = "123456"
)

type counterCache_struct struct {
	List sync.Map //alis
}

var (
	CounterCache = newCache() //单例
)

func newCache() *counterCache_struct {
	//return cache.NewCache("key", `{"key":" `+REDIS_COLLECTIONNAME+`","conn":":`+REDIS_PORT+`","dbNum":"0","password":"`+REDIS_PASSWORLD+`"}`)
	return &counterCache_struct{
		List: sync.Map{},
	}
}

//向计数器插入值
//@uid:计时器编号
//@number:计时器值，支持负数  范围限制 +10 ~ -10
func (c *counterCache_struct) Add(uid string, number string) error {

	//获取或插入
	cnt, ok := c.List.LoadOrStore(uid, Counter{
		Uid:     uid,
		Count:   0,
		Expired: "3600",
	})

	if ok {
		k := cnt.(Counter)
		curNum, err := strconv.ParseInt(number, 10, 64)
		if err != nil {
			return err
		}

		k.Count = k.Count + curNum

		c.List.Store(uid, k)
		return nil
	}

	if cnt == nil {
		return errors.New("Invalid uid")
	}
	return nil
}

//删除计数器
//@uid:计数器id
func (c *counterCache_struct) Del(uid string) error {
	c.List.Delete(uid)
	return nil
}

//查询计数器
//@uid:计数器id
//@return: 计数值
func (c *counterCache_struct) Get(uid string) (error, string) {
	v, ok := c.List.Load(uid)
	if !ok {
		return errors.New("No data"), ""
	}
	return nil, fmt.Sprint(v.(Counter).Count)
}

//输出全部计数器
func (c *counterCache_struct) ShowList() (error, []Counter) {
	var cnts []Counter
	c.List.Range(func(key, value interface{}) bool {
		switch value.(type) {
		case Counter:
			cnts = append(cnts, value.(Counter))
			return true
		default:
			return false
		}
	})

	if len(cnts) == 0 {
		return errors.New("No datas"), nil
	}

	return nil, cnts
}
