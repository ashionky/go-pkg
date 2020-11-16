/**
 * @Author pibing
 * @create 2020/11/14 1:22 PM
 */

package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-pkg/errcode"
	"go-pkg/params"
	"go-pkg/pkg/redis"
	"go-pkg/svc"
	"go-pkg/util"
	"net/http"
	"strings"
	"time"
)

// 鉴权处理器，pri的接口都需要鉴权才能访问
func Authorize(c *gin.Context) {
	token := c.GetHeader("X-Token")
	if token == "" {
		// 没有登陆过
		c.AbortWithStatusJSON(http.StatusOK, errcode.Resp(errcode.NoToken))
		return
	}

	tokenData, err := redis.Get(util.FormatTokenUserKey(token))
	if err != nil {
		if strings.Contains(err.Error(), "redigo: nil returned") {
			c.AbortWithStatusJSON(http.StatusOK, errcode.Resp(errcode.InvalidToken))
		} else {
			c.AbortWithStatusJSON(http.StatusOK, errcode.Resp(errcode.InternalError))
		}
		return
	}

	ut := util.UserToken{}
	err = json.Unmarshal(tokenData, &ut)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, errcode.Resp(errcode.InvalidToken))
		return
	}

	if ut.UID == 0 || ut.ExpireTime < time.Now().Unix() {
		c.AbortWithStatusJSON(http.StatusOK, errcode.Resp(errcode.TokenExpired))
		return
	}

	// 在header中追加用户id
	c.Request.Header.Set("X-UID", fmt.Sprint(ut.UID))
	c.Next()
}


//  用户登录
func SignIn(c *gin.Context) {
	var param params.SigninReq
	err := c.ShouldBind(&param)
	if err != nil {
		c.JSON(http.StatusOK, errcode.Resp(errcode.InvalidParams, err.Error()))
		return
	}
	rsp, ae := svc.SignIn(&param)
	c.JSON(http.StatusOK, errcode.Resp(ae, rsp))
}

//登出
func SignOut(c *gin.Context) {
	ae := svc.SignOut(getToken(c))
	c.JSON(http.StatusOK, errcode.Resp(ae))
}

func getToken(c *gin.Context) string {
	return c.Request.Header.Get("X-Token")
}

