package main

import (
	"errors"
	"sort"
	"sync"
)

type MapList struct {
	rw sync.RWMutex
	mq map[string](*[]TodoList)
}

func NewMapList() *MapList {
	return &MapList{
		rw: sync.RWMutex{},
		mq: make(map[string](*[]TodoList)),
	}
}

func (c *MapList) AddTodoList(pid string, todoList *TodoList) {
	c.rw.Lock()
	defer c.rw.Unlock()
	lst, ok := c.mq[pid]
	if !ok {
		l := &[]TodoList{*todoList}
		c.mq[pid] = l
		return
	}
	/*
		1.插入排序
	*/

	for i := 0; i < len(*lst); i++ {
		//若存在 则更新
		if (*lst)[i].Sid == todoList.Sid {
			(*lst)[i] = *todoList
			c.mq[pid] = lst
			return
		}
	}

	*lst = append(*lst, *todoList)
	//此处可优化 改由插排
	sort.Sort(TodoListArray(*lst))
	c.mq[pid] = lst
	return
}

func (c *MapList) GetTodoList(pid string) (todoLists *[]TodoList, err error) {
	c.rw.RLock()
	defer c.rw.RUnlock()

	lst, ok := c.mq[pid]
	if !ok {
		return nil, errors.New("No data")
	}

	return lst, nil
}

func (c *MapList) DelTodoList() {
}
