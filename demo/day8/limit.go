package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func init()  {
	rand.Seed(time.Now().Unix())
}
func main()  {
	timer := make(chan struct{})
	channel := make(chan interface{})
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		time.Sleep(time.Second * 5)
		timer <- struct{}{}
		wg.Done()
	}()

	go func() {
		s := rand.Intn(10)
		fmt.Println(s)
		time.Sleep(time.Duration(s)*time.Second)
		channel <- s
		wg.Done()
	}()

	select {
	case e:=<-timer:
		fmt.Println("time out : ",e)
	case e:=<-channel:
		fmt.Println("output : ",e)
	}
}
