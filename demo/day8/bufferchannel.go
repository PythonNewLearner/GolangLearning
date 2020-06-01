package main

import (
	"fmt"
	"sync"
)

func main()  {
	channel := make(chan int,3)
	go func() {
		for i:=0;i<10;i++{
			channel <- i
		}
		close(channel)   // dont forget close the channel
	}()

	var wg sync.WaitGroup
	wg.Add(1)    // 1 个历程
	go func() {       //第二个历程会自动阻塞（自动等待第一个历程读取数据）（第二个历程依赖于第一个历程）
		for e:=range channel{
			fmt.Printf("%d\n",e)
		}
		wg.Done()
	}()


	wg.Wait()
}
