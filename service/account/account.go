package account

import (
	"younghe/lib/errorx"
	"younghe/lib/mysqlx"
	"younghe/model/account"
)

func LoginAccount(email, password string) (*account.Account, error) {
	val := &account.Account{}
	existed, err := mysqlx.DB.Where("email = ? and password = ?", email, password).Get(val)
	if err != nil {
		return nil, err
	}
	if !existed {
		return nil, errorx.NewAuthError("账号或密码错误!")
	}
	return val, err
}

func CreateAccount(email, password string) error {
	val := &account.Account{Email: email, Password: password}
	_, err := mysqlx.DB.InsertOne(val)
	return err
}
