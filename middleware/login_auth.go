package middleware

import (
	"context"
	"github.com/go-redis/redis/v8"
	"younghe/config"
	"younghe/lib/corex"
	"younghe/lib/errorx"
	"younghe/lib/redisx"
)

func LoginAuth() corex.HandlerFunc {
	return func(c *corex.Context) {
		sessionId, err := c.Cookie(config.SessionCookieName)
		if err != nil || sessionId == "" {
			c.XErrorResponse(errorx.NewAuthError("请登录后再操作!"))
			return
		}

		sessionInfoBytes, err := redisx.Redis.Get(context.TODO(), config.SessionRedisPrefix+sessionId).Bytes()
		if err != nil {
			if err == redis.Nil {
				c.XErrorResponse(errorx.NewAuthError("请登录后再操作!"))
				return
			}
			c.XErrorResponse(err)
			return
		}
		c.Set(config.SessionContextName, sessionInfoBytes)
		c.Next()
	}
}
