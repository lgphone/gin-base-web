package account

import (
	"context"
	"encoding/json"
	"time"
	"younghe/config"
	"younghe/lib/common"
	"younghe/lib/corex"
	"younghe/lib/errorx"
	"younghe/lib/redisx"
	account_service "younghe/service/account"
)

func Login(c *corex.Context) {
	req := struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.XErrorResponse(errorx.NewParamError(err.Error()))
		return
	}
	accountInfo, err := account_service.LoginAccount(req.Email, req.Password)
	if err != nil {
		c.XErrorResponse(err)
		return
	}
	sessionId := common.GetRandomUUID()
	sessionInfo := &corex.SessionInfo{
		Username: accountInfo.Email,
		Id:       100,
	}
	sessionInfoBytes, err := json.Marshal(sessionInfo)
	if err != nil {
		c.XErrorResponse(err)
		return
	}
	c.SetCookie(config.SessionCookieName, sessionId, config.Config.SessionTimeout, "/", "", false, true)
	if err = redisx.Redis.Set(context.TODO(), config.SessionRedisPrefix+sessionId, sessionInfoBytes, time.Duration(config.Config.SessionTimeout)*time.Second).Err(); err != nil {
		c.XErrorResponse(err)
		return
	}
	c.XSuccessResponse(sessionInfo)
}

func GetOwn(c *corex.Context) {
	sessionInfo, err := c.XGetSessionInfo()
	if err != nil {
		// log err
		c.XErrorResponse(err)
		return
	}
	// session时间续期
	sessionId, _ := c.Cookie(config.SessionCookieName)
	c.SetCookie(config.SessionCookieName, sessionId, config.Config.SessionTimeout, "/", "", false, true)
	if err = redisx.Redis.Expire(context.TODO(), config.SessionRedisPrefix+sessionId, time.Duration(config.Config.SessionTimeout)*time.Second).Err(); err != nil {
		c.XErrorResponse(err)
		return
	}
	c.XSuccessResponse(sessionInfo)
}

func Logout(c *corex.Context) {
	sessionId, _ := c.Cookie(config.SessionCookieName)
	if sessionId != "" {
		c.SetCookie(config.SessionCookieName, sessionId, -1, "/", "", false, true)
		if err := redisx.Redis.Expire(context.TODO(), config.SessionRedisPrefix+sessionId, 0).Err(); err != nil {
			c.XErrorResponse(err)
			return
		}
	}
	c.XSuccessResponse("退出成功!")
}

func Register(c *corex.Context) {
	req := struct {
		Email     string `json:"email" binding:"required"`
		Password  string `json:"password" binding:"required"`
		Password2 string `json:"password2" binding:"required"`
	}{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.XErrorResponse(errorx.NewParamError(err.Error()))
		return
	}
	if req.Password != req.Password2 {
		c.XErrorResponse(errorx.NewParamError("两次输入的密码不一致!"))
		return
	}
	if err := account_service.CreateAccount(req.Email, req.Password); err != nil {
		c.XErrorResponse(err)
		return
	}
	c.XSuccessResponse("注册成功!")
}
