package main

import (
	"fmt"
	"time"
)

type User struct{
	id int
	name string
	addr string
	tel string
}

type Task struct {
	id int
	name string
	startTime time.Time
	endTime time.Time
	status int
	*User   //匿名嵌入
}

func main()  {
	var task Task
	fmt.Printf("%#v\n",task)

	task = Task{
		id: 1,
		name: "DoneTask",
		User : &User{
			id:2,
			name:"Jason",
			addr:"USA",
		},
	}
	fmt.Printf("%#v\n",task)
	fmt.Println(
		task.name,
		task.User.name,
		task.User.addr,
		task.addr,
	)


}