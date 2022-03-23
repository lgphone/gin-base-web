package demo

import "time"

type Demo struct {
	Id        int64
	Name      string    `xorm:"varchar(25) notnull unique comment('姓名')"`
	Age       uint      `xorm:"notnull"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
	DeletedAt time.Time `xorm:"deleted"`
}
