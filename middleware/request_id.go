package middleware

import (
	"younghe/lib/common"
	"younghe/lib/corex"
)

func RequestId() corex.HandlerFunc {
	return func(c *corex.Context) {
		requestId := c.Request.Header.Get("X-Request-Id")
		if requestId == "" {
			requestId = common.GetRandomUUID()
		}
		c.Set("X-Request-Id", requestId)
		c.Writer.Header().Set("X-Request-Id", requestId)
		c.Next()
	}
}
