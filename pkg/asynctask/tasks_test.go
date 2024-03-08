package asynctask

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestNewTask(t *testing.T) {
	fmt.Println("==================================TestNewTask=================================")
	ctx := context.Background()
	handler := func(s string, ctx context.Context) { //ctx context.Context必填，放到末尾，是继承Run函数中的ctx
		fmt.Println("aaaaaaa", s, ctx)
		return
	}
	task := NewTask(handler, "aaaaaaaaabbbbbbbaaaaaaaaaaaaa")
	task.Run(ctx)
	//task.Wait()

	fmt.Println(task.Result)
}

func TestWaitAll(t *testing.T) {
	fmt.Println("==================================TestWaitAll=================================")
	handler1 := func(ctx context.Context) { //ctx context.Context必填，放到末尾，是继承Run函数中的ctx
		fmt.Println(time.Now())
		fmt.Println("handler1", "我在等待指定的时间后执行")
	}
	param2 := "aaaaaaaaaaaaaaaaaaaaaa"
	handler2 := func(p string, ctx context.Context) { //ctx context.Context必填，放到末尾，是继承Run函数中的ctx
		fmt.Println("handler2", time.Now())
		fmt.Println(p)
	}
	param3 := "bbbbbbbbbbbbbbbbbbbbbbbbbb"
	handler3 := func(p string, ctx context.Context) string { //ctx context.Context必填，放到末尾，是继承Run函数中的ctx
		fmt.Println(p)
		return p + "111111111111111"
	}
	ctx := context.Background()
	task1 := NewTask(handler1).ContinueWith(func(result TaskResult) {
		fmt.Println("我在task1执行后执行。")
	}).ContinueWith(func(result TaskResult) {
		fmt.Println("我在task1执行后第二次执行。")
	}).Delay(5 * time.Second).Run(ctx)
	task2 := NewTask(handler2, param2).Run(ctx)
	task3 := NewTask(handler3, param3).Run(ctx)

	WaitAll(task1, task2, task3)
	fmt.Println(task3.Result)
}
