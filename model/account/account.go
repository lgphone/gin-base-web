package account

import "time"

type Account struct {
	Id        int64
	Email     string    `xorm:"varchar(25) notnull unique comment('邮箱')"`
	Password  string    `xorm:"notnull"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
	DeletedAt time.Time `xorm:"deleted"`
}
