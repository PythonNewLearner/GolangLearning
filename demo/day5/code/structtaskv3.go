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
type Creator struct {
	id int
	name string
	addr string
}
type Task struct {
	id int
	name string
	startTime time.Time
	endTime time.Time
	status int
	User   // anony embedded : 面向对象的继承;这个也有属性名，属性名是 User
	//Creator   // Creator和User有相同的属性，这时候要用全名访问
}

func main()  {
	var task Task
	fmt.Printf("%#v\n",task)
	fmt.Printf("%#v\n",task.User.name)

	task.User.name = "chen"

	fmt.Printf("%#v\n",task)

	//赋值
	task = Task {
		id: 1,
		name: "Done",
		User: User{
			id: 1,
			name: "Peter",
			addr: "Singapore",
		},
	}
	fmt.Printf("%#v\n",task)

	fmt.Println(task.name)
	fmt.Println(task.addr)

	task.addr = "HongKong"
	fmt.Println(task.addr)
	fmt.Printf("%#v\n",task)

	task.User.addr = "USA"
	fmt.Println(task.addr)
	fmt.Printf("%#v\n",task)


	

}