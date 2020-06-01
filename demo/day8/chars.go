package main

import (
	"fmt"
	"time"
)

func chars(prefix string)  {
	for i:= 'A';i<='Z';i++{
		fmt.Printf("%s %c\n",prefix,i)
		time.Sleep(time.Microsecond)
	}
	
}
func main()  {
	go chars("goroutine")
	chars("main")

}
