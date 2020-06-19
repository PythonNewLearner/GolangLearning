package main

import (
	"fmt"
	"sync"
)

func f()  {
	fmt.Println("hello")
}

func main()  {
	var once sync.Once
	for i:=0;i<10;i++{
		once.Do(f)
	}
}