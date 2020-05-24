package main

import (
	"fmt"
	"strconv"
)

var todos = []map[string]string{
	{"id":"1","name":"读书","startTime":"18:00","endTime":"","status":statusNew,"user":"chen"},
	{"id":"2","name":"复习","startTime":"20:00","endTime":"","status":statusNew,"user":"chen"},
	{"id":"3","name":"发呆","startTime":"21:00","endTime":"","status":statusNew,"user":"chen"},
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
func main()  {

	var text string
	task := newTask()
	fmt.Println("请输入任务信息：")

	fmt.Print("任务名：")
	fmt.Scan(&text)
	task[name] = text

	fmt.Print("开始时间：")
	fmt.Scan(&text)
	task[startTime] = text

	fmt.Print("负责人")
	fmt.Scan(&text)
	task[user] = text

	todos = append(todos,task)
	fmt.Println(todos)
}