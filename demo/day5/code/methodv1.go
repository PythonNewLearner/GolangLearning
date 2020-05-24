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

//user
func (task *Task)SetUser(user string)  {   // task Task 是值传递，要传递指针才能改变属性值
	task.user = user
}

func main()  {

	//方法值： 实例.方法名
	task1 := Task{}     //值
	task2 := &Task{}   //指针

	method1 := task1.SetName
	method2 := task2.SetName   // 自动解引用
	fmt.Printf("%#v\n",method1)
	fmt.Printf("%#v\n",method2)
	method1("Todo1")     //改不了
	method2("Todo2")     //改不了
	fmt.Printf("%#v\n",task1)
	fmt.Printf("%#v\n",task2)


	method3 := task1.SetUser
	method4 := task2.SetUser 
	method3("Kate")
	method4("Lucy")
	fmt.Printf("%#v\n",task1)    //能改
	fmt.Printf("%#v\n",task2)    //能改
	//方法表达式: 结构体.方法名

	
}