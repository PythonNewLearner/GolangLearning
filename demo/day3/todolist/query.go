package main

import (
	"fmt"

	"strings"
)
//
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
func printTask(task map[string]string)  {
	fmt.Println(strings.Repeat("-",20))
	fmt.Println("id:",task[id])
	fmt.Println("任务名称:",task[name])
	fmt.Println("开始时间:",task[startTime])
	fmt.Println("结束时间:",task[endTime])
	fmt.Println("状态:",task[status])
	fmt.Println("负责人:",task[user])
}
func main()  {
	var text string
	fmt.Println("请输入查询信息:")
	fmt.Scan(&text)
	for _,todo :=range todos{
		if strings.Contains(todo[name],text){
			printTask(todo)
		}
	}
}