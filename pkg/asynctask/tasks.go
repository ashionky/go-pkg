package asynctask

import (
	"container/list"
	"context"
	"fmt"
	"reflect"
	"runtime/debug"
	"sync"
	"time"
)

//参数
type TaskParameter interface{}

//执行的方法
type TaskHandler interface{}

//等待任务执行完成的后续任务
type ContinueWithHandler func(TaskResult)

//返回的参数类型
type TaskResult struct {
	Result interface{}
	Error  error
}

//一个任务
type Task struct {
	wait         *sync.WaitGroup
	handler      reflect.Value
	params       []reflect.Value
	Result       TaskResult //任务执行完成的返回结果
	once         sync.Once
	IsCompleted  bool //表示任务是否执行完成
	continueWith *list.List
	delay        time.Duration
}

//新建一个任务
func NewTask(handler TaskHandler, params ...TaskParameter) *Task {

	handlerValue := reflect.ValueOf(handler)

	if handlerValue.Kind() == reflect.Func {
		task := Task{
			wait:         &sync.WaitGroup{},
			handler:      handlerValue,
			IsCompleted:  false,
			continueWith: list.New(),
			delay:        0 * time.Second,
			params:       make([]reflect.Value, 0),
		}
		if paramNum := len(params); paramNum > 0 {
			task.params = make([]reflect.Value, paramNum)
			for index, v := range params {
				task.params[index] = reflect.ValueOf(v)
			}
		}
		return &task
	}
	panic("handler not func")
}

//启动异步任务
func (task *Task) Run(parentCtx context.Context) *Task {
	ctx := context.Background()
	ctx = AddTrace(parentCtx, ctx)

	task.once.Do(func() {
		task.wait.Add(1)
		go func(ctx interface{}) {
			if task.delay.Nanoseconds() > 0 {
				time.Sleep(task.delay)
			}
			defer func() {
				task.IsCompleted = true
				if task.continueWith != nil {
					result := task.Result
					for element := task.continueWith.Back(); element != nil; element = element.Prev() {
						if tt, ok := element.Value.(ContinueWithHandler); ok {
							tt(result)
						}
					}
				}
				if panicInfo := recover(); panicInfo != nil {
					fmt.Println("async panic Info", panicInfo, string(debug.Stack()))
				}
				task.wait.Done()
				ctx.(context.Context).Done() //执行完毕，销毁context
			}()
			task.params = append(task.params, reflect.ValueOf(ctx))
			values := task.handler.Call(task.params)

			task.Result = TaskResult{
				Result: values,
			}
		}(ctx)
	})

	return task
}

//等待任务完成
func (task *Task) Wait() {
	task.wait.Wait()
}

//等待所有任务都完成
func WaitAll(tasks ...*Task) {
	wait := &sync.WaitGroup{}
	for _, task := range tasks {
		wait.Add(1)
		go func(task *Task) {
			defer wait.Done()
			task.wait.Wait()
		}(task)
	}
	wait.Wait()
}

//立即启动一个异步任务
func StartNew(ctx context.Context, handler TaskHandler, params ...TaskParameter) *Task {
	var task *Task
	if len(params) == 0 {
		task = NewTask(handler)
	} else {
		task = NewTask(handler, params...)
	}
	task.Run(ctx)
	return task
}

//当前Task执行完后执行
func (task *Task) ContinueWith(handler ContinueWithHandler) *Task {
	task.continueWith.PushFront(handler)

	return task
}

//延迟指定的时间后执行
func (task *Task) Delay(delay time.Duration) *Task {
	task.delay = delay
	return task
}
