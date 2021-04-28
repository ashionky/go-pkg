/**
 * @Author pibing
 * @create 2021/4/26 9:04 AM
 */

package ratelimit

import (
	"fmt"
	"testing"
	"time"
)

func TestM(t *testing.T)  {
	LimitHandler(5,1, func() {
		fmt.Println(time.Now())
	})
}
