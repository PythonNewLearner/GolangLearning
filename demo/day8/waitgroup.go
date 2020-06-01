package main

import (
	"fmt"
	"sync"
	"time"
)
func Mainchars(prefix string)  {
	for i:= 'A';i<='Z';i++{
		fmt.Printf("%s %c\n",prefix,i)
		time.Sleep(time.Microsecond)
	}

}
func chars(prefix string,wg *sync.WaitGroup)  {
	for i:= 'A';i<='Z';i++{
		fmt.Printf("%s %c\n",prefix,i)
		time.Sleep(time.Microsecond)
	}
	wg.Done()

}
func main()  {
	var wg sync.WaitGroup
	wg.Add(1)
	go chars("goroutine",&wg)
	Mainchars("main")
	wg.Wait()
	fmt.Println("over")
}

