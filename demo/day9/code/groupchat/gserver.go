package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"sync"
)

func main()  {
	addr := "0.0.0.0:9888"
	listener,err := net.Listen("tcp",addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	var wg sync.WaitGroup
	//var locker sync.Mutex



		for {
			wg.Add(1)
			go func() {
				conn,err :=listener.Accept()
				if err != nil {
					log.Println(err)
				}
				fmt.Println(conn.RemoteAddr()," connected")

				func(){
					//defer func() {
					//	fmt.Println(conn.RemoteAddr(),"disconnected")
					//	conn.Close()
					//}()
					ctx := make([]byte,1024*1024)
					//locker.Lock()
					reader := bufio.NewReader(conn)
					reader.Read(ctx)

					fmt.Println(conn.RemoteAddr(),"says:",string(ctx))
					//locker.Unlock()
				}()
			}()
			wg.Done()
			}


	wg.Wait()
}
