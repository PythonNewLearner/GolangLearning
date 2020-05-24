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
	user User   // embedded
}

func main()  {
	var task Task
	fmt.Printf("%#v\n",task)

	task.user.name = "chen"
	fmt.Printf("%#v\n",task)
	
	task = Task{
		id:1,
		name:"DONE",
		user: User{
			id : 1,
			name : "baichen",
		},
	}
	fmt.Println(task)

	user := User{id:2,name:"jason"}
	task = Task{
		id:1,
		name:"done",
		user: user,
	}
	fmt.Println(task)

	task.user = User{id:3,name:"Peter"}
	fmt.Println(task)
}