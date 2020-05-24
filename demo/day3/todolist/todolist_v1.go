package main

import (
	"fmt"
	"strconv"
	"strings"
)

//1. 任务的输入（添加任务）
//2. 任务列表（任务查询）
//3.任务修改
//4.任务删除
//5. 详情

//任务名称，开始时间，结束时间，状态，负责人
// id,name,startTime,endTime,status,user
// []map[string]string
var todos = []map[string]string{
	{"id":"1","name":"读书","startTime":"18:00","endTime":"","status":statusNew,"user":"chen"},
	{"id":"2","name":"复习","startTime":"20:00","endTime":"","status":statusNew,"user":"chen"},
	{"id":"3","name":"发呆","startTime":"21:00","endTime":"","status":statusNew,"user":"chen"},
	{"id":"4","name":"练功","startTime":"11:00","endTime":"","status":statusNew,"user":"chen"},
	{"id":"5","name":"上山","startTime":"21:00","endTime":"","status":statusNew,"user":"chen"},
	{"id":"6","name":"买东西","startTime":"21:00","endTime":"","status":statusNew,"user":"chen"},
	{"id":"7","name":"吃饭","startTime":"21:00","endTime":"","status":statusNew,"user":"chen"},
	{"id":"8","name":"买课本","startTime":"21:00","endTime":"","status":statusNew,"user":"chen"},
}
const  (
	id = "id"
	name = "name"
	startTime = "startTime"
	endTime = "endTime"
	status = "status"
	user = "user"
)
const (
	statusNew = "新创建"
	statusComplete = "完成"
	statusIncomplete = "未完成"

)

func genId() int {
	var rt int
	for _,todo :=range todos{
		todoid,_ := strconv.Atoi(todo["id"])
		if  todoid > rt{
			rt = todoid
		}
	}
	return rt + 1
}
func printTask(task map[string]string)  {
	fmt.Println(strings.Repeat("-",20))
	fmt.Println("id:",task[id])
	fmt.Println("任务名称:",task[name])
	fmt.Println("开始时间:",task[startTime])
	fmt.Println("结束时间:",task[endTime])
	fmt.Println("状态:",task[status])
	fmt.Println("负责人:",task[user])
	fmt.Println(strings.Repeat("-",20))
}
func newTask() map[string]string  {
	task := make(map[string]string)
	task[id] = strconv.Itoa(genId())
	task[name] = ""
	task[startTime] = ""
	task[endTime] = ""
	task[status] = statusNew
	task[user] = ""
	return task
}
func add()  {
	task := newTask()
	fmt.Println("请输入任务信息：")

	task[name] = input("任务名")
	task[startTime] = input("开始时间")
	task[user] = input("负责人")
	todos = append(todos,task)
	fmt.Println("创建任务成功！")
	//fmt.Println(todos)
}
func query()  {
	var text string
	text = input("请输入查询信息(all 查询所有任务)：")
	for _,todo :=range todos{
		if text =="all" || strings.Contains(todo[name],text){
			printTask(todo)
		}
	}
}
func input(prompt string) string  {
	var text string
	fmt.Println(prompt)
	fmt.Scan(&text)
	return text

}
func main()  {
	for {
		text := input("请输入操作:")
		switch text {
		case "add":
			add()
		case "query":
			query()
		case "modify":
		case "delete":
		case "exit":
			return
		default:
			fmt.Println("输入请求不正确")
		}

	}
}
