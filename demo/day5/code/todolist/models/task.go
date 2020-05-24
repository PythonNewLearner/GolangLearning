package models

import (
	"time"
)

//属性名首字母也要大写
type Task1 struct{
	id int
	name string
	startTime *time.Time
	endTime *time.Time
	user string
}

type Task2 struct{
	Id int
	Name string
	StartTime *time.Time
	EndTime *time.Time
	User string
}

type task3 struct{  //不能被访问
	Id int
	Name string
	StartTime *time.Time
	EndTime *time.Time
	User string
}

type task4 struct{
	Id int
	Name string
	StartTime *time.Time
	EndTime *time.Time
	User string
}

type Anonytask struct{
	task4
}

