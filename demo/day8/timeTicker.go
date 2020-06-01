package main

import (
	"fmt"
	"time"
)

func main()  {
	ticker := time.NewTicker(time.Second * 3)
	defer ticker.Stop()

	end := time.Now().Add(time.Second*20)
	for t := range ticker.C {
		fmt.Println(t)
		if t.After(end){
			break
		}
	}
}
