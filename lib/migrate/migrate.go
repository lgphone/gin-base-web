package migrate

import (
	"younghe/lib/mysqlx"
	"younghe/model/account"
	"younghe/model/demo"
)

func Migrate() {
	if err := mysqlx.DB.Sync2(&demo.Demo{}, &account.Account{}); err != nil {
		panic(err)
	}
}
