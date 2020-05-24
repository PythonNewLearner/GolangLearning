package main
import (
	"fmt"
	"time"
)
type Task struct {
	id int
	name string
	startTime time.Time
	endTime time.Time
	status int
	user string
}
func main()  {
	var task Task
	fmt.Printf("%T\n",task)
	fmt.Printf("%#v\n",task)

	//assignment 

	task = Task{}
	fmt.Printf("%#v\n",task) //  0

	task = Task{1,"homework",time.Now(),time.Now().Add(24*time.Hour),1,"chen"}   // assign value
	fmt.Printf("%#v\n",task)

	task = Task{
		id:2,
		name:"shopping",
		user:"chen",
		}    // assign value
	fmt.Printf("%#v\n",task)


	// pointer
	var ptask *Task
	fmt.Printf("%T\n",ptask)
	fmt.Printf("%#v\n",ptask)

	//assignment
	ptask = &Task{}
	fmt.Printf("%#v\n",ptask)


	fmt.Println(task.id,task.status,task.name)
	task.status = 3
	fmt.Println(task.status)

	//值类型
		task1 := task
		task1.user = "bai"
		fmt.Println(task)
		fmt.Println(task1)
		change_task(task)   //不变
		fmt.Println(task)


		//初始化结构体（指针）方式
		ptask2 := new(Task)
		fmt.Printf("%T\n",ptask2)
		fmt.Printf("%#v\n",ptask2)

		// 引用类型





	
}

func change_task(task Task)  {
	task.user = "xxx"
}