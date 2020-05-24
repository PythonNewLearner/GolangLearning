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
	var task1 Task
	fmt.Printf("%#v\n",task1)

	task1 = Task{
		id: 1,
		name: "DoneTask",
		User : &User{
			id:2,
			name:"Jason",
			addr:"USA",
		},
	}
	task2 := task1
	task2.name = "notDone"
	fmt.Println(task1)
	fmt.Println(task2)

	task1.User.addr = "Beijing"      //指针类型对task1和task2都有影响
	fmt.Printf("%#v\n",task1.User)
	fmt.Printf("%#v\n",task2.User)



}