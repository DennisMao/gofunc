package main

import (
	"errors"
	"sort"
	"sync"
)

type MapMap struct {
	rw sync.RWMutex
	mq map[string](map[string]*TodoList)
}

func NewMapMap() *MapMap {
	return &MapMap{
		rw: sync.RWMutex{},
		mq: make(map[string](map[string]*TodoList)),
	}
}
func (c *MapMap) AddTodoList(pid string, todoList *TodoList) {
	c.rw.Lock()
	defer c.rw.Unlock()

	lst, ok := c.mq[pid]
	if !ok {
		l := make(map[string]*TodoList)
		c.mq[pid] = l
		return
	}

	//直接更新
	lst[todoList.Sid] = todoList
	return

}

func (c *MapMap) GetTodoList(pid string) (todoLists *[]TodoList, err error) {
	c.rw.RLock()
	defer c.rw.RUnlock()

	lst, ok := c.mq[pid]
	if !ok {
		return nil, errors.New("No datas")
	}

	var ret []TodoList
	for _, v := range lst {
		ret = append(ret, *v)
	}

	/*
		1.按时间排序
	*/
	sort.Sort(TodoListArray(ret))

	return &ret, nil
}

func (c *MapMap) DelTodoList() {
}
