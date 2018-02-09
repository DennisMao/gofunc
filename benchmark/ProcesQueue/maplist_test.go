package main

import (
	"fmt"
	"testing"
	"time"
)

////////////////////// 功能测试 //////////////////////
// 写入测试
func TestMapList(t *testing.T) {
	t.Log("map list rw test")
	lmq := NewMapList()

	tdList := make([]TodoList, 100)
	tnow := time.Now()
	for i := 0; i < 100; i++ {
		tdList[i].Sid = fmt.Sprintf("%d", i)
		tdList[i].Status = "processing"
		tdList[i].NoticeTime = tnow.Add(1 * time.Second).Format("2006-01-02 15:04:05")
	}

	t.Log("insert sid=1~10000 ")
	bg := time.Now()
	for i := 0; i < 100; i++ {
		lmq.AddTodoList("admin", &tdList[i])
	}
	t.Log("time used = " + time.Now().Sub(bg).String())
}

//	读取测试
func TestMapListGet(t *testing.T) {

	//初始化
	lmq := collecttMapList1K()

	//读取
	lst, err := lmq.GetTodoList("admin_500")
	if err != nil {
		t.Log(err)
		return
	}
	for i := 0; i < len(*lst); i++ {
		t.Log((*lst)[i])
	}
	return
}

////////////////////// 性能测试 //////////////////////
func BenchmarkMapListInsert(b *testing.B) {
	b.StopTimer()

	//初始化
	lmq := NewMapList()

	//事件数量 100条
	tdList := make([]TodoList, 100)
	tnow := time.Now()
	for i := 0; i < 100; i++ {
		tnow = tnow.Add(1 * time.Second)
		tdList[i].Sid = fmt.Sprintf("%d", i)
		tdList[i].Status = "processing"
		tdList[i].NoticeTime = tnow.Format("2006-01-02 15:04:05")
	}

	b.StartTimer()

	//添加事件
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(tdList); j++ {
			lmq.AddTodoList(fmt.Sprintf("admin_%d", i), &tdList[j])
		}
	}
}

//	读取测试
func BenchmarkMapListGetList1K(b *testing.B) {
	b.StopTimer()

	//初始化数据
	lmq := collecttMapList1K()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		lmq.GetTodoList(fmt.Sprintf("admin_%d", i))
	}
}

//	读取测试
func BenchmarkMapListGetList1M(b *testing.B) {
	b.StopTimer()

	//初始化数据
	lmq := collectMapList1M()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		lmq.GetTodoList(fmt.Sprintf("admin_%d", i))
	}
}

//	测试数据准备
//	1K Queues x 100 elements
func collecttMapList1K() *MapList {
	//初始化
	lmq := NewMapList()

	//事件数量 100条
	tdList := make([]TodoList, 100)
	tnow := time.Now()
	for i := 0; i < 100; i++ {
		tdList[i].Sid = fmt.Sprintf("%d", i)
		tdList[i].Status = "processing"
		tdList[i].NoticeTime = tnow.Add(1 * time.Second).Format("2006-01-02 15:04:05")
	}

	//添加事件
	for i := 0; i < 1000; i++ {
		for j := 0; j < len(tdList); j++ {
			lmq.AddTodoList(fmt.Sprintf("admin_%d", i), &tdList[j])
		}
	}

	return lmq
}

//	测试数据准备
//	1M Queues x 100 elements
func collectMapList1M() *MapList {
	//初始化
	lmq := NewMapList()

	//事件数量 100条
	tdList := make([]TodoList, 100)
	tnow := time.Now()
	for i := 0; i < 100; i++ {
		tdList[i].Sid = fmt.Sprintf("%d", i)
		tdList[i].Status = "processing"
		tdList[i].NoticeTime = tnow.Add(1 * time.Second).Format("2006-01-02 15:04:05")
	}

	//添加事件
	for i := 0; i < 100000; i++ {
		for j := 0; j < len(tdList); j++ {
			lmq.AddTodoList(fmt.Sprintf("admin_%d", i), &tdList[j])
		}
	}

	return lmq
}
