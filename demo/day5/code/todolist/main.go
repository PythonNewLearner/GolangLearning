package main

import (
	"fmt"
	"todolist/models"
)

func main()  {	
	var task  models.Task2 = models.Task2{
		Id:1,
		Name:"washing",
	}
	fmt.Printf("%#v\n",task)

	task.User = "chen"
	fmt.Printf("%#v\n",task)

	task4  := models.Anonytask {}
	fmt.Printf("%#v\n",task4)

	task4.Name = "John"
	fmt.Printf("%#v\n",task4)
}