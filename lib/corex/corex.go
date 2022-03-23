package corex

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"runtime/debug"
	"strconv"
	"younghe/config"
	"younghe/lib/errorx"
	"younghe/lib/loggerx"
)

type SessionInfo struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
}

type Context struct {
	*gin.Context
}

type respStruct struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  interface{} `json:"msg"`
}

func (c *Context) XSuccessResponse(data interface{}) {
	c.JSON(200, &respStruct{
		Code: 1000,
		Data: data,
		Msg:  nil,
	})
	return
}

func (c *Context) XErrorResponse(e error) {
	resp := &respStruct{Data: nil}
	switch ev := e.(type) {
	case *errorx.AuthError:
		resp.Code = ev.Code
		resp.Msg = ev.Error()
	case *errorx.ParamError:
		resp.Code = ev.Code
		resp.Msg = ev.Error()
	case *errorx.LogicError:
		resp.Code = ev.Code
		resp.Msg = ev.Error()
	default:
		loggerx.Logger.Error(fmt.Printf("%s", debug.Stack()))
		resp.Code = 1005
		resp.Msg = ev.Error()
	}
	c.JSON(200, resp)
	c.Abort()
}

func (c *Context) XHasUserId() bool {
	if c.Query("user_id") != "" {
		return true
	}
	return false
}

func (c *Context) XGetSessionInfo() (*SessionInfo, error) {
	rawData, ok := c.Get(config.SessionContextName)
	if !ok {
		return nil, errorx.NewLogicError("没有session数据! ")
	}
	sessionInfo := &SessionInfo{}
	byteData, _ := rawData.([]byte)
	if err := json.Unmarshal(byteData, &sessionInfo); err != nil {
		return nil, err
	}
	return sessionInfo, nil
}

func (c *Context) XPageParams() (int, int) {
	page := c.Query("page")
	pageSize := c.Query("page_size")
	if page != "" && pageSize != "" {
		pVal, err1 := strconv.Atoi(page)
		psVal, err2 := strconv.Atoi(pageSize)
		if err1 != nil || err2 != nil {
			return 1, 15
		}
		return pVal, psVal
	}
	return 1, 15
}

type HandlerFunc func(c *Context)

func Handle(h HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := &Context{
			c,
		}
		h(ctx)
	}
}
