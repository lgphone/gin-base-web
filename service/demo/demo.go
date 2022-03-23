package demo

import (
	"errors"
	"fmt"
	"younghe/lib/mysqlx"
	"younghe/model/demo"
)

func GetOneWithId(id uint) (*demo.Demo, error) {
	val := &demo.Demo{}
	existed, err := mysqlx.DB.ID(id).Get(val)
	if !existed {
		return val, errors.New(fmt.Sprintf("没有查询到id为%v的数据!", id))
	}
	return val, err
}

func GetAll() ([]*demo.Demo, error) {
	val := make([]*demo.Demo, 0)
	err := mysqlx.DB.Find(&val)
	return val, err
}

func GetListWithPage(page, pageSize int) ([]*demo.Demo, int64, error) {
	val := make([]*demo.Demo, 0)
	var count int64
	err := mysqlx.DB.Limit(pageSize, (page-1)*pageSize).Find(&val)
	if err != nil {
		return val, count, err
	}
	count, err = mysqlx.DB.Count(&demo.Demo{})
	return val, count, err
}

func DeleteOneWithId(id uint) error {
	deleted, err := mysqlx.DB.ID(id).Delete(&demo.Demo{})
	if deleted == 0 {
		return errors.New("资源未找到,删除失败! ")
	}
	return err
}

func CreateOne(val *demo.Demo) error {
	_, err := mysqlx.DB.InsertOne(val)
	return err
}
