package concurrency

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

// Note:
// 任务目标
// 从字符串数组中统计出现'a'的个数

// 顺序执行
//
// 协程数: 1
func Single(raw []string, subStr string) int {
	cnt := int(0)

	for i, _ := range raw {
		cnt += strings.Count(raw[i], subStr)
	}

	return cnt
}

// 单条并行执行
//
// 协程数: len(raw)
// 共享方式: 加锁
func ConcurrencyByLock(raw []string, subStr string) int {
	cnt := int(0)
	cntRw := new(sync.RWMutex)
	wg := new(sync.WaitGroup)

	for i, _ := range raw {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()

			// 业务
			c := strings.Count(raw[i], subStr)

			// 收集&整理结果
			cntRw.Lock()
			cnt += c
			cntRw.Unlock()
		}(wg)
	}

	wg.Wait()
	return cnt
}

// 单条并行执行
//
// 协程数: len(raw)
// 共享方式: channel
func ConcurrencyByChannel(raw []string, subStr string) int {
	cnt := int(0)
	cntChan := make(chan int, len(raw))
	wg := new(sync.WaitGroup)

	for i, _ := range raw {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()

			// 业务
			c := strings.Count(raw[i], subStr)

			// 收集结果
			cntChan <- c
		}(wg)
	}

	wg.Wait()
	close(cntChan)

	// 整理结果
	for c := range cntChan {
		cnt += c
	}

	return cnt
}

// 单条并行执行,带终止,外部可紧急终止任务
//
// 协程数: len(raw)
// 共享方式:  channel
// 异常处理: 只要有一条任务出现异常返回失败
func ConcurrencyByChannelWithCancelControl(ctx context.Context, raw []string, subStr string) int {
	cnt := int(0)
	cntChan := make(chan int, len(raw))
	wg := new(sync.WaitGroup)

	for i, _ := range raw {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ctx context.Context) {
			defer wg.Done()

			select {
			case <-ctx.Done():
				return
			default:
			}

			// 业务
			c := strings.Count(raw[i], subStr)

			// 收集结果
			cntChan <- c

		}(wg, ctx)

	}

	wg.Wait()
	close(cntChan)

	// 整理结果
	for c := range cntChan {
		cnt += c
	}

	return cnt
}

// 单条并行执行,带超时,外部可设置任务的超时时间
//
// 协程数: len(raw)
// 共享方式:  channel
// 异常处理: 只要有一条任务出现异常返回失败
func ConcurrencyByChannelWithTimeoutControl(raw []string, subStr string, timeout time.Duration) int {
	cnt := int(0)
	cntChan := make(chan int, len(raw))
	wg := new(sync.WaitGroup)
	ctx, _ := context.WithTimeout(context.Background(), timeout)

	for i, _ := range raw {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ctx context.Context) {
			defer wg.Done()

			select {
			case <-ctx.Done():
				return
			default:
			}

			// 业务
			c := strings.Count(raw[i], subStr)

			// 收集结果
			cntChan <- c

		}(wg, ctx)

	}

	wg.Wait()
	close(cntChan)

	// 整理结果
	for c := range cntChan {
		cnt += c
	}

	return cnt
}

// 单条并行执行,带错误处理,遇到错误提示失败
//
// 协程数: len(raw)
// 共享方式:  channel
// 异常处理: 只要有一条任务出现异常返回失败
func ConcurrencyByChannelWithErrorControl(raw []string, subStr string) (int, error) {
	cnt := int(0)
	cntChan := make(chan int, len(raw))
	wg := new(errgroup.Group)

	for i, _ := range raw {
		wg.Go(func() error {

			// 业务
			c := strings.Count(raw[i], subStr)

			// 收集结果
			cntChan <- c

			return fmt.Errorf("test_error")
		})

	}

	err := wg.Wait()
	close(cntChan)
	if err != nil {
		return 0, err
	}

	// 整理结果
	for c := range cntChan {
		cnt += c
	}

	return cnt, nil
}

// 切割并行执行
//
// 协程数: len(raw) / splitGranularity
// 共享方式:  channel
func SplitConcurrency(raw []string, subStr string, splitGranularity int) int {
	if splitGranularity == 0 {
		splitGranularity = 1
	}

	cnt := int(0)
	cntChan := make(chan int, len(raw))
	lenRaw := len(raw)
	part := lenRaw / splitGranularity
	if lenRaw%2 != 0 && lenRaw > 0 {
		part += 1
	}
	wg := new(sync.WaitGroup)

	for i := 0; i < part; i++ {
		st := i * splitGranularity
		end := (i + 1) * splitGranularity
		if end > lenRaw {
			end = lenRaw
		}

		wg.Add(1)
		go func(wg *sync.WaitGroup, raw []string, subStr string) {
			defer wg.Done()

			for i, _ := range raw {
				// 业务
				c := strings.Count(raw[i], subStr)

				// 收集结果
				cntChan <- c
			}

		}(wg, raw[st:end], subStr)
	}

	wg.Wait()
	close(cntChan)

	// 整理结果
	for c := range cntChan {
		cnt += c
	}

	return cnt
}

// 协程池
//
// 协程数: 自定义
// 共享方式: channel
func ConcurrencyByGoroutinePool() {
	//TODO
	//...
}
