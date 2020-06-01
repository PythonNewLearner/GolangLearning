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
	var a,b = 10000,10000
	var wg sync.WaitGroup
	var locker sync.Mutex
		wg.Add(2)
		go func() {
			for i:=0;i<100;i++{
				amt := rand.Intn(100)
				func () {
					locker.Lock()
					defer locker.Unlock()
					a += amt
					time.Sleep(time.Microsecond)
					b -= amt
				}()
			}
			wg.Done()
		}()

		go func() {
			for i:=0;i<100;i++{
				amt := rand.Intn(100)
				func() {
					locker.Lock()
					defer locker.Unlock()
					a -= amt
					time.Sleep(time.Microsecond)
					b += amt
				}()
			}
			wg.Done()
		}()

	wg.Wait()
	fmt.Println(a,b,a+b)
}
