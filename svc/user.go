/**
 * @Author pibing
 * @create 2020/11/14 1:26 PM
 */

package svc

import (
	"encoding/json"
	"go-pkg/errcode"
	"go-pkg/model"
	"go-pkg/params"
	"go-pkg/pkg/db"
	"go-pkg/pkg/redis"
	"go-pkg/util"
	"gorm.io/gorm"
	"time"
)

// 用户登陆
func SignIn(param *params.SigninReq) (rsp *params.SigninRsp, ae errcode.APIError) {
	var user model.User
	err := db.GetDB().Model(&model.User{}).Where(nil).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errcode.UserNotExists
		}
		return nil, errcode.MysqlFailed
	}

	if user.Password != param.Password {
		return nil, errcode.UsernameOrPasswordError
	}

	// 生成token
	rsp = &params.SigninRsp{}
	rsp.UID = user.ID
	rsp.Token = util.GenToken()

	ut := util.UserToken{
		Name:       user.Name,
		UID:        user.ID,
		Token:      rsp.Token,
		ExpireTime: time.Now().Unix() + util.TokenExpireTime,
	}

	tokenData, err := json.Marshal(ut)
	if err != nil {
		return nil, errcode.InternalError
	}

	// 把以前的token干掉
	oldToken, _ := redis.Get(util.FormatUserTokenKey(rsp.UID))
	if oldToken != nil {
		redis.Delete(util.FormatTokenUserKey(string(oldToken)))
	}

	// 记录uid和token的互相关联关系
	redis.Set(util.FormatUserTokenKey(rsp.UID), rsp.Token)
	redis.Set(util.FormatTokenUserKey(rsp.Token), tokenData)

	// todo 把用户信息写到缓存，可提高访问效率

	return rsp, errcode.Success
}


// 用户登出
func SignOut(token string) errcode.APIError {
	tokenData, err := redis.Get(util.FormatTokenUserKey(token))
	if err != nil {
		return errcode.RedisError
	}
	redis.Delete(util.FormatTokenUserKey(token))

	ut := util.UserToken{}
	err = json.Unmarshal(tokenData, &ut)
	if err != nil {
		return errcode.InternalError
	}
	redis.Delete(util.FormatUserTokenKey(ut.UID))

	return errcode.Success
}
