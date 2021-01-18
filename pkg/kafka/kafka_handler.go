/**
 * @Author pibing
 * @create 2021/1/16 4:49 PM
 */

package kafka

import (
	"context"
	"sync"
)

//注册消息处理handler
func InitHandler(ctx context.Context, w *sync.WaitGroup) {

	testEventHandler := TestEventHandler{}
	for i := 0; i < 2; i++ {
		w.Add(1)
		go NewConsumerHandler("test", "test", ctx, w, testEventHandler)
	}

	//其它topic消息处理-----

}
