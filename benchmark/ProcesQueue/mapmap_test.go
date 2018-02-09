package main

import (
	"fmt"
	"testing"
	"time"
)

////////////////////// 功能测试 //////////////////////
//	写入测试
func TestMapMapInsert(t *testing.T) {

	//初始化
	lmq := NewMapMap()

	//事件数量 100条
	tdList := make([]TodoList, 100)
	tnow := time.Now()
	for i := 0; i < 100; i++ {
		tnow = tnow.Add(1 * time.Second)
		tdList[i].Sid = fmt.Sprintf("%d", i)
		tdList[i].Status = "processing"
		tdList[i].NoticeTime = tnow.Format("2006-01-02 15:04:05")
	}

	//添加事件
	for i := 0; i < 100000; i++ {
		for j := 0; j < len(tdList); j++ {
			lmq.AddTodoList(fmt.Sprintf("admin_%d", i), &tdList[j])
		}
	}
}

//	读取测试
func TestMapMapGet(t *testing.T) {

	//初始化
	lmq := collectMapMap1K()

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
//	写入测试
func BenchmarkMapMapInsert(b *testing.B) {
	b.StopTimer()

	//初始化
	lmq := NewMapMap()

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
func BenchmarkMapMapGetList1K(b *testing.B) {
	b.StopTimer()

	//初始化
	lmq := collectMapMap1K()

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		lmq.GetTodoList(fmt.Sprintf("admin_%d", i))
	}

}
func BenchmarkMapMapGetList1M(b *testing.B) {
	b.StopTimer()

	//初始化
	lmq := collectMapMap1M()

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		lmq.GetTodoList(fmt.Sprintf("admin_%d", i))
	}

}

////////////////////// 测试数据准备 //////////////////////
//	1K Queues x 100 elements
func collectMapMap1K() *MapMap {
	//初始化
	lmq := NewMapMap()

	//事件数量 100条
	tdList := make([]TodoList, 100)
	tnow := time.Now()
	for i := 0; i < 100; i++ {
		tnow = tnow.Add(1 * time.Second)
		tdList[i].Sid = fmt.Sprintf("%d", i)
		tdList[i].Status = "processing"
		tdList[i].NoticeTime = tnow.Format("2006-01-02 15:04:05")
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
func collectMapMap1M() *MapMap {
	//初始化
	lmq := NewMapMap()

	//事件数量 100条
	tdList := make([]TodoList, 100)
	tnow := time.Now()
	for i := 0; i < 100; i++ {
		tnow = tnow.Add(1 * time.Second)
		tdList[i].Sid = fmt.Sprintf("%d", i)
		tdList[i].Status = "processing"
		tdList[i].NoticeTime = tnow.Format("2006-01-02 15:04:05")
	}

	//添加事件
	for i := 0; i < 100000; i++ {
		for j := 0; j < len(tdList); j++ {
			lmq.AddTodoList(fmt.Sprintf("admin_%d", i), &tdList[j])
		}
	}

	return lmq
}
