package asynctask

import (
	"context"
	"fmt"
	"testing"
)

func TestNewSTask(t *testing.T) {
	fmt.Println("==================================TestNewSTask=================================")
	ctx := context.Background()
	s := "abc"
	handler := func(ctx context.Context) {
		fmt.Println(s)
		return
	}
	task := Go(ctx, handler)
	task.Wait()
}

func TestRegex(t *testing.T) {

}
