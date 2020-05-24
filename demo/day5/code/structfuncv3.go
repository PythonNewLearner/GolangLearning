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

func NewTask(id int,name string,user string) *Task  {
	return &Task{
		id:id,
		name:name,
		startTime:time.Now()
		user:user,
	}
}

//定义方法 指定类型：接受者

//name
func (task Task)SetName(name string)  {   
	task.name = name
}

//user
func (task *Task)SetUser(user string)  {   // task Task 是值传递，要传递指针才能改变属性值
	task.user = user
}

func main()  {
	task1 := Task{}
	task1.SetName("Done")

	task1.SetUser("baichen")    //  这是go的语法糖 相当于(&task).SetUser("baichen") 
	fmt.Printf("%#v\n",task1)   

	task2 := &Task{name:"OK"}  
	task2.SetName("Not")   //这也是语法糖，但是不会改变值  相当于(*task2).SetName("Not")
	fmt.Println(task2)
	//总结以上两种，跟两种定义是有关系的

	//指针是nil是不能调用方法的

	
}