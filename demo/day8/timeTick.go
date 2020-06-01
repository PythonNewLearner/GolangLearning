package main

import (
	"fmt"
	"time"
)

func main()  {

	stop :=time.Now().Add(time.Second*20)
	for t:=range time.Tick(time.Second*3){
		fmt.Println(t)
		if t.After(stop){
			break
		}
	}
}
