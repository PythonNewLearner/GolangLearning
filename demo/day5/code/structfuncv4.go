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

func NewUser(id int,name ,addr ,tel string) *User {
	return &User{id:id,name:name,addr:addr,tel:tel}
}
func (u *User) SetName(name string)  {
	u.name = name
}

func (u *User) SetAddr(addr string)  {
	u.addr = addr
}

type Task struct {
	id int
	name string
	startTime *time.Time
	endTime *time.Time
	status int
	*User   // embedded
}
func NewTask(id int,name string,startTime, endTime *time.Time,user *User) *Task  {
	return &Task{id,name,startTime,endTime,1,user}
}
func main()  {
	start := time.Now()
	end :=  start.Add(24*time.Hour)
	user := NewUser(1,"chen","Singapore","123456")
	task :=  NewTask(1,"Done",&start,&end,user)
	fmt.Printf("%#v\n",task)
	fmt.Printf("%#v\n",task.User)

	task.SetName("John")  // task.User.SetName
	task.SetAddr("USA")		// task.User.SetAddr
	fmt.Printf("%#v\n",task)
	fmt.Printf("%#v\n",task.User)
	
	
}