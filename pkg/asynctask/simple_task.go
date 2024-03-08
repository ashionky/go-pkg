package asynctask

import (
	"context"
	"fmt"
	"runtime/debug"
	"sync"
)

/**
 * @Author: Charlie
 * @Email: zhouchunli@meross.com
 * @Date: 2021/3/18 14:52
 * @Desc:
 */

//一个任务
type STask struct {
	wait    *sync.WaitGroup
	handler func(ctx context.Context)
	once    sync.Once
}

//新建一个任务
func NewSTask(handler func(ctx context.Context)) *STask {
	STask := STask{
		wait:    &sync.WaitGroup{},
		handler: handler,
	}
	return &STask
}

//启动异步任务
func (st *STask) Run(parentCtx context.Context) *STask {
	ctx := context.Background()
	ctx = AddTrace(parentCtx, ctx)
	st.once.Do(func() {
		st.wait.Add(1)
		go func(ctx context.Context) {
			defer func() {
				if panicInfo := recover(); panicInfo != nil {
					fmt.Println("async panic Info", panicInfo, string(debug.Stack()))
				}
				st.wait.Done()
				ctx.Done() //执行完毕，销毁context
			}()
			st.handler(ctx)
		}(ctx)
	})
	return st
}

//等待任务完成
func (st *STask) Wait() {
	st.wait.Wait()
}

//立即启动一个异步任务
func Go(ctx context.Context, handler func(ctx context.Context)) *STask {
	var st *STask
	st = NewSTask(handler).Run(ctx)
	return st
}
