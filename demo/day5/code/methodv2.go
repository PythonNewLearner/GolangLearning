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

//name
func (task Task)SetName(name string)  {   
	task.name = name
}

/* 自动生成  :  method2 := (*Task).SetName
func (task *Task)SetName(name string)  {   
	(*task).name = name
}
*/

//user
func (task *Task)SetUser(user string)  {   // task Task 是值传递，要传递指针才能改变属性值
	task.user = user
}

func main()  {

	//方法表达式: 结构体.方法名
	task1 := Task{}     //值
	task2 := &Task{}   //指针
	fmt.Printf("%#v\n",task1)
	fmt.Printf("%#v\n",task2)

	//对于值接受者，可以通过指针/值来获取方法表达式
	method1 := Task.SetName
	method1(task1,"test")
	method1(*task2,"test")
	fmt.Printf("%#v\n",task1)
	fmt.Printf("%#v\n",task2)

	method2 := (*Task).SetName
	method2(&task1,"test")
	method2(task2,"test")
	fmt.Printf("%#v\n",task1)
	fmt.Printf("%#v\n",task2)

	fmt.Printf("%#v\n",method1)  // func(main.Task, string))
	fmt.Printf("%#v\n",method2)   // func(*main.Task, string))

	//对于指针接受者，只能通过指针来获取方法表达式
	//method3 := Task.SetUser
	method4 := (*Task).SetUser
	fmt.Printf("%#v\n",method4)    //  func(*main.Task, string))
	method4(&task1,"test")
	method4(task2,"test")
	fmt.Printf("%#v\n",task1)
	fmt.Printf("%#v\n",task2)



	
}