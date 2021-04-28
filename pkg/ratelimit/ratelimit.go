/**
 * @Author pibing
 * @create 2021/4/25 5:47 PM
 */

package ratelimit

import (
	"go.uber.org/ratelimit"
	"time"
)

//限速 某个时间内执行多少次函数调用
//per 时间
//rate 执行次数
//调用的函数
func LimitHandler(per,rate int, handler func())  {
	time_per := ratelimit.Per(time.Duration(per) * time.Second)
	rl := ratelimit.New(rate,time_per)
	count:=1
	for  {
		rl.Take()
		handler()
		count++
	}

}
