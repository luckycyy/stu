package model

import "time"

//{"id":"0001","name":"张三","namesx":"zs","grade":"5","class":"1","tel":"138","arrive":"false","arrivetime":"11"},
type Student struct {
	Id int
	Name string
	Namesx string
	Grade int
	Class string
	Tel string
	Arrive bool
	Arrivetime time.Time
}