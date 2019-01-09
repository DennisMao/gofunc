ync.Cond的使用 

在多协程同时监听消费单个变量时候,我们通常会采用channel这种go的无锁通信来进行,但对于一些已有的结构体,我们还是需要进行加读写锁(RWmutex)共享的方式来消费。典型应用如下代码所示:
```
type Data struct {
	Value Queue
	Rw *sync.RWMutex
}

func Consume(data *Data){
	data.Rw.RLock()
	defer data.Rw.RUnlock()
	
	value := data.Value.Pop() //消费
}

func Produce(data *Data){
	data.Rw.Lock()
	defer data.Rw.Unlock()
	
	data.Value.Push(XX)  //生产
}
```
通常的内存共享的过程:
+ 1.生成者对变量加锁 rw.Lock()
+ 2.生产者生产
+ 3.生成者完成生成解锁 rw.Unlock()
+ 4.消费者常驻等待  rw.RLock()
+ 5.消费者消费
+ 6.消费者完成消费  rw.RUnlock()


当生产者对写锁释放后消费者会立即获得读锁进行消费,这种情况生成者生成完后无法把控消费者的消费情况,比如:  
+ 1.阻塞消费者,等生成完成后再唤醒消费消费者
+ 2.生成完成后,控制消费的延迟开始
+ 3.只让一个消费者消费

对于上述情况`sync.Cond`包则提供了解决方法。  

```
type Data struct {
	Value Queue
	rw *sync.RWMutex
	Cond *sync.Cond
}

func Consume(data *Data){
	data.Rw.Lock()
	defer data.Rw.Unlock()
	
	data.Cond.Wait() //阻塞等待生成者的唤醒
	
	value := data.Value.Pop() //消费
}

func Produce(data *Data){
	data.Rw.Lock()
	defer data.Rw.Unlock()
	
	data.Value.Push(XX)  //生产
	
	//通知一个消费者消费
	//data.Cond.Signal()
	//通知全部消费者消费
	data.Cond.Broadcast()
}
```

实际测试代码:
```
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
```
输出结果:
>2019/01/06 11:49:28 生产者: 生产1
2019/01/06 11:49:28 消费者:1  启动
2019/01/06 11:49:28 消费者:2  启动
2019/01/06 11:49:29 生产者: 生产2
2019/01/06 11:49:29 消费者:1  消费:2
2019/01/06 11:49:29 消费者:2  消费:2
2019/01/06 11:49:30 生产者: 生产3
2019/01/06 11:49:30 消费者:1  消费:3
2019/01/06 11:49:30 消费者:2  消费:3


引用:
+ [sync.Cond](https://godoc.org/sync#Cond)

