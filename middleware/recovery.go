package middleware

import (
	"net"
	"net/http/httputil"
	"os"
	"strings"
	"younghe/lib/corex"
	"younghe/lib/loggerx"
)

func Recovery() corex.HandlerFunc {
	return func(c *corex.Context) {
		defer func() {
			if err := recover(); err != nil {
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}
				loggerx.Logger.Error(string(httpRequest), err)
				c.XErrorResponse(err.(error))
				return
			}
		}()
		c.Next()
	}
}
