# Go并发模型

写了一个示例代码,讲解一些实际工作中用到的并发模型(生产者/消费者组合)和一些并发下的函数处理方案.

在使用go并发时候有以下内容需要注意:
+ 1.在协程的函数中注意return,防止协程泄露 (泄露可以通过pprof看出)
+ 2.协程内的panic会引起整个程序的panic,但go语言并不支持跨协程去recover,所以根据业务需要在协程内的函数做好recover处理。
+ 3.协程入参建议加入'context',方便传参和控制
+ 4.注意`race`情况的发生,debug时候使用[go test -race]检测。

## 无并发
Single  
[代码](https://github.com/DennisMao/goexperience/blob/370cebbf079e58342b3da7e96ac011fdbaa307d1/concurrency/concurrent.go#L20)
```
func Single(raw []string, subStr string) int
```
## 单任务并发
### 共享内存: 锁
通过`sync.Mutex`、`sync.RWMutex`或`sync.atomi`锁机制实现对内存进行协程间的共享  
[代码](https://github.com/DennisMao/goexperience/blob/370cebbf079e58342b3da7e96ac011fdbaa307d1/concurrency/concurrent.go#L34)
```
func ConcurrencyByLock(raw []string, subStr string) int
```
### 共享内存: channel
通过`channel`实现对内存进行协程间的共享  
[代码](https://github.com/DennisMao/goexperience/blob/370cebbf079e58342b3da7e96ac011fdbaa307d1/concurrency/concurrent.go#L62)
```
func ConcurrencyByChannel(raw []string, subStr string) int
```

### 带任务终止控制
外部引入一个context,可使用context.WithCancel生成的context来对子函数进行控制。父级context.Done时候,全部子级都会返回。
这种设计要求在子协程里都必须加上context的控制。当然,用chan信号量也是可以实现。    
[代码](https://github.com/DennisMao/goexperience/blob/370cebbf079e58342b3da7e96ac011fdbaa307d1/concurrency/concurrent.go#L96)
```
func ConcurrencyByChannelWithCancelControl(ctx context.Context, raw []string, subStr string) int 
```

### 带任务超时控制
内部引入一个context,可使用context.WithTimeout生成的context来对子函数进行控制。[tips: 用context.WithDeadline()同样可以实现,根据业务场景使用,如果是需要相对时间的用WithTimeout会方便一些,需要精确时间点的用Deadline]父级context.Done时候,全部子级都会返回。
这种设计要求在子协程里都必须加上context的控制。  
[代码](https://github.com/DennisMao/goexperience/blob/370cebbf079e58342b3da7e96ac011fdbaa307d1/concurrency/concurrent.go#L139)
```
func ConcurrencyByChannelWithTimeoutControl(raw []string, subStr string, timeout time.Duration) int
```

### 带异常处理
引入`errgroup`来替换原有的`waitGroup`,只要有一条任务失败,整个任务失败返回  
[代码](https://github.com/DennisMao/goexperience/blob/370cebbf079e58342b3da7e96ac011fdbaa307d1/concurrency/concurrent.go#L183)
```
func ConcurrencyByChannelWithErrorControl(raw []string, subStr string) (int, error)
```

## 切割任务并发
### 共享内存: channel
将原有数组数据按一定方法切割成多个任务再执行多个协程去处理,转化为 `多生产者`模型。这个思想有点`map-reduce`的感觉。
这个`map`的过程可以根据原始数据的特性来进行处理。切割任务的模型也是平常业务中用到比较多的模型,根据实际场景把大的数据按一定规则切割成合理个
小数据并发执行,最后合并。  
[代码](https://github.com/DennisMao/goexperience/blob/370cebbf079e58342b3da7e96ac011fdbaa307d1/concurrency/concurrent.go#L220)
```
func SplitConcurrency(raw []string, subStr string, splitGranularity int) int
```

#### 生产者消费者并行处理
在共享内存channel的基础上,实现消费者与生产者并行执行,以加速整体任务的执行效率。


## 协程池
协程池的思想近似于把`方法2:共享内存`的协程中的channel指针给保存到一个`sync.Pool`、Map或者一个新的channel里。方法2的协程变成`常驻`协程。当有业务传参来的时候,分配一个已启动的常驻协程的入口channel暴露给外部,传入数据进行处理。这样能大大减小协程的反复开启销毁的开销。


# 引用
+ [Go内存模型](https://golang.org/ref/mem)
+ [Go竞争探测器](https://golang.org/doc/articles/race_detector.html)
+ [通过通信来共享内存](https://golang.org/doc/codewalk/sharemem/)
+ [Go并发模型](https://www.youtube.com/watch?v=f6kdp27TYZs)
