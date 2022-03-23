package demo

import (
	"context"
	"younghe/lib/convert"
	"younghe/lib/corex"
	"younghe/lib/errorx"
	"younghe/lib/redisx"
	"younghe/model/demo"
	demo_service "younghe/service/demo"
)

func Ping(c *corex.Context) {
	c.XSuccessResponse("pong")
}

func GetAll(c *corex.Context) {
	val, err := demo_service.GetAll()
	if err != nil {
		c.XErrorResponse(err)
		return
	}
	resp := make([]*Resp, 0)
	for _, item := range val {
		resp = append(resp, &Resp{
			Id:        item.Id,
			Name:      item.Name,
			Age:       item.Age,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		})
	}
	c.XSuccessResponse(resp)
}

func PageList(c *corex.Context) {
	val, count, err := demo_service.GetListWithPage(c.XPageParams())
	if err != nil {
		c.XErrorResponse(err)
		return
	}

	respList := &RespList{
		Count: count,
		List:  make([]*Resp, 0),
	}
	for _, item := range val {
		respList.List = append(respList.List, &Resp{
			Id:        item.Id,
			Name:      item.Name,
			Age:       item.Age,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		})
	}
	c.XSuccessResponse(respList)
}

func CreateOne(c *corex.Context) {
	req := struct {
		Age  uint   `json:"age" binding:"required"`
		Name string `json:"name" binding:"required"`
	}{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.XErrorResponse(errorx.NewParamError(err.Error()))
		return
	}
	val := &demo.Demo{Name: req.Name, Age: req.Age}
	if err := demo_service.CreateOne(val); err != nil {
		c.XErrorResponse(err)
		return
	}
	c.XSuccessResponse("创建成功!")
}

func GetOne(c *corex.Context) {
	Id := convert.StrTo(c.Query("id")).MustUInt()
	if Id == 0 {
		c.XErrorResponse(errorx.NewParamError("id 不能为空!"))
		return
	}
	val, err := demo_service.GetOneWithId(Id)
	if err != nil {
		c.XErrorResponse(err)
		return
	}
	resp := &Resp{
		Id:        val.Id,
		Name:      val.Name,
		Age:       val.Age,
		CreatedAt: val.CreatedAt,
		UpdatedAt: val.UpdatedAt,
	}
	c.XSuccessResponse(resp)
}

func DeleteOne(c *corex.Context) {
	Id := convert.StrTo(c.Query("id")).MustUInt()
	if Id == 0 {
		c.XErrorResponse(errorx.NewParamError("id 不能为空!"))
		return
	}
	if err := demo_service.DeleteOneWithId(Id); err != nil {
		c.XErrorResponse(err)
		return
	}
	c.XSuccessResponse("删除成功!")
}

func RedisGet(c *corex.Context) {
	redisKey := c.Query("key")
	val, err := redisx.Redis.Get(context.TODO(), redisKey).Result()
	if err != nil {
		c.XErrorResponse(err)
		return
	}
	c.XSuccessResponse(val)
}

func RedisSet(c *corex.Context) {
	req := struct {
		RedisKey   string `json:"redis_key" binding:"required"`
		RedisValue string `json:"redis_value" binding:"required"`
	}{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.XErrorResponse(errorx.NewParamError(err.Error()))
		return
	}
	val, err := redisx.Redis.Set(context.TODO(), req.RedisKey, req.RedisValue, 0).Result()
	if err != nil {
		c.XErrorResponse(err)
		return
	}
	c.XSuccessResponse(val)
}

func PanicTest(c *corex.Context) {
	list := []string{"1", "2"}
	c.XSuccessResponse(list[3])
}
