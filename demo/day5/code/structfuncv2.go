package main
import (
	"fmt"
	"time"
)

type Task struct{
	id int
	name string
	startTime *time.Time
	endTime *time.Time
	user string
}

//定义方法 指定类型：接受者
func (task *Task)SetName(name string)  {   // task Task 是值传递，要传递指针才能改变属性值
	task.name = name
}
func (task *Task)GetName() string  {
	return task.name
}
func main()  {
	task := &Task{}
	task.SetName("Done")
	fmt.Println(task)

	task.SetName("TODO")
	fmt.Println(task.GetName())
	
}