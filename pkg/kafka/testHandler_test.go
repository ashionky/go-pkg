/**
 * @Author pibing
 * @create 2020/11/15 12:54 PM
 */

package kafka

import (
	"context"
	"sync"
	"testing"
)

func TestConsumer(t *testing.T)  {
	var ctx context.Context
	var w *sync.WaitGroup
	testEventHandler := TestEventHandler{}
	for i := 0; i < 2; i++ {
		w.Add(1)
		go NewConsumerHandler("test", "test", ctx, w, testEventHandler)
	}
	return
}
