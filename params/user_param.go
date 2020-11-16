package params

// 登陆请求
type SigninReq struct {
	// 手机号
	Phone string `json:"phone" form:"phone" binding:"required"`
	// 用户密码
	Password string `json:"password" form:"password" binding:"required"`
}

// 登陆响应
type SigninRsp struct {
	// 用户id
	UID uint `json:"uid" form:"uid" binding:"required"`
	// 用户令牌，后续请求都需要在header X-Token带上此token
	Token string `json:"token" form:"token" binding:"required"`
}


