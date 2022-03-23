package demo

import "time"

type Resp struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	Age       uint      `json:"age"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RespList struct {
	Count int64   `json:"count"`
	List  []*Resp `json:"list"`
}
