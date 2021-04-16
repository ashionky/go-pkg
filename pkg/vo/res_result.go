/**
 * @Author pibing
 * @create 2020/11/18 1:49 PM
 */

package vo

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Result struct {
	Code int         `json:"code"` //1：成功   其它code码请自行定义
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

/*
	构造响应对象
*/
func GetDefaultResult() *Result {
	return &Result{1, "", nil}
}

func (res *Result) Set(code int, msg string) {
	res.Code = code
	res.Msg = msg
}
func (res *Result) SetMsg(msg string) {
	res.Msg = msg
}

/**
请求处理成功，发送响应的统一处理
*/
func SendSuccess(ctx *gin.Context, rep *Result) {
	rep.Set(1, rep.Msg)
	ctx.AbortWithStatusJSON(http.StatusOK, rep)
}

/**
请求处理失败，发送响应的统一处理
*/
func SendFailure(ctx *gin.Context, code int, rep *Result) {
	rep.Set(code, rep.Msg)
	ctx.AbortWithStatusJSON(http.StatusOK, rep)
}
