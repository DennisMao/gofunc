package main

import (
	"sync"
	"time"
)

import (
	"log"
)

type Data struct {
	Value []int
	rw    *sync.RWMutex
	Cond  *sync.Cond
}

func main() {
	data := Data{make([]int, 0), &sync.RWMutex{}, nil}
	data.Cond = sync.NewCond(data.rw)

	wg := &sync.WaitGroup{}

	producer := func() {
		defer wg.Done()

		for {
			//生产过程
			data.Cond.L.Lock()
			newValue := data.Value[len(data.Value)-1] + 1
			log.Printf("生产者: 生产%d\n", newValue)
			data.Value = append(data.Value, newValue)
			data.Cond.L.Unlock()

			//通知一个消费者消费
			//data.Cond.Signal()
			//通知全部消费者消费
			data.Cond.Broadcast()

			time.Sleep(1 * time.Second)
		}
	}

	consumer := func(idx int) {
		defer wg.Done()

		log.Printf("消费者:%d  启动\n", idx)

		for {
			//消费过程
			data.Cond.L.Lock()
			data.Cond.Wait() //阻塞,等待消费信号
			newValue := data.Value[len(data.Value)-1]
			data.Cond.L.Unlock()

			log.Printf("消费者:%d  消费:%d\n", idx, newValue)

		}
	}

	wg.Add(3)
	go consumer(1)
	go consumer(2)
	go producer()

	wg.Wait()
}

