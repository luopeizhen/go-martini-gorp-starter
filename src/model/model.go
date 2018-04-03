package model

type Table1 struct {
	Id   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}
