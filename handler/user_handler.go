/**
 * @Author pibing
 * @create 2020/11/14 1:22 PM
 */

package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-pkg/constant"
	"go-pkg/logic"
	"go-pkg/params"
	redis "go-pkg/pkg/go-redis"
	"go-pkg/pkg/vo"
	"go-pkg/util"
	"strings"
	"time"
)

// 鉴权处理器，pri的接口都需要鉴权才能访问
func Authorize(c *gin.Context) {
	res := vo.GetDefaultResult()
	token := c.GetHeader("X-Token")
	if token == "" {
		// 没有登陆过
		vo.SendFailure(c, constant.NoToken, res)
		return
	}

	tokenData, err := redis.Get(util.FormatTokenUserKey(token))

	if err != nil {
		if strings.Contains(err.Error(), "redis: nil") {
			vo.SendFailure(c, constant.InvalidToken, res)
		} else {
			vo.SendFailure(c, constant.InternalError, res)
		}
		return
	}

	ut := util.UserToken{}
	err = json.Unmarshal([]byte(tokenData), &ut)
	if err != nil {
		vo.SendFailure(c, constant.InvalidToken, res)
		return
	}

	if ut.UID == 0 || ut.ExpireTime < time.Now().Unix() {
		vo.SendFailure(c, constant.TokenExpired, res)
		return
	}

	// 在header中追加用户id
	c.Request.Header.Set("X-UID", fmt.Sprint(ut.UID))
	c.Next()
}

//  用户登录
func SignIn(c *gin.Context) {
	res := vo.GetDefaultResult()
	var param params.SigninReq
	err := c.ShouldBind(&param)
	if err != nil {
		vo.SendFailure(c, constant.InvalidParams, res)
		return
	}
	rsp, err := logic.SignIn(&param)
	if err != nil {
		vo.SendFailure(c, constant.InternalError, res)
		return
	}
	res.Data = rsp
	vo.SendSuccess(c, res)
}

//登出
func SignOut(c *gin.Context) {
	res := vo.GetDefaultResult()
	err := logic.SignOut(getToken(c))
	if err != nil {
		vo.SendFailure(c, constant.InternalError, res)
		return
	}
	vo.SendSuccess(c, res)
}

func getToken(c *gin.Context) string {
	return c.Request.Header.Get("X-Token")
}
