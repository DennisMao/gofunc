# Go并发模型

写了一个示例代码,讲解一些实际工作中用到的并发模型(生产者/消费者组合)和一些并发下的函数处理方案.

## 无并发
Single

## 单任务并发
### 共享内存: 锁
```
func ConcurrencyByLock(raw []string, subStr string) int
```
### 共享内存: channel
```
func ConcurrencyByChannel(raw []string, subStr string) int
```

### 带任务终止控制
外部引入一个context,可使用context.WithCancel生成的context来对子函数进行控制。父级context.Done时候,全部子级都会返回。
这种设计要求在子协程里都必须加上context的控制。当然,用chan信号量也是可以实现。

```
func ConcurrencyByChannelWithCancelControl(ctx context.Context, raw []string, subStr string) int 
```

### 带任务超时控制
内部引入一个context,可使用context.WithTimeout生成的context来对子函数进行控制。[tips: 用context.WithDeadline()同样可以实现,根据业务场景使用,如果是需要相对时间的用WithTimeout会方便一些,需要精确时间点的用Deadline]父级context.Done时候,全部子级都会返回。
这种设计要求在子协程里都必须加上context的控制。
```
func ConcurrencyByChannelWithTimeoutControl(raw []string, subStr string, timeout time.Duration) int
```

### 带异常处理
引入`errgroup`来替换原有的`waitGroup`,只要有一条任务失败,整个任务失败返回
```
func ConcurrencyByChannelWithErrorControl(raw []string, subStr string) (int, error)
```

## 切割任务并发
### 共享内存: channel
将原有数组数据按一定方法切割成多个任务再执行多个协程去处理,转化为 `多生产者`模型。这个思想有点`map-reduce`的感觉。
这个`map`的过程可以根据原始数据的特性来进行处理。切割任务的模型也是平常业务中用到比较多的模型,根据实际场景把大的数据按一定规则切割成合理个
小数据并发执行,最后合并。
```
func SplitConcurrency(raw []string, subStr string, splitGranularity int) int
```
## 多生产者多消费者模型
TODO
### 使用类令牌桶控制生产者/消费者数量
