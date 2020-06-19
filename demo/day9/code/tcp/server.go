package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main()  {
	listener,err :=net.Listen("tcp","127.0.0.1:9999")
	if err != nil{
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn,err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
			func(){
				defer conn.Close()
				log.Println(conn.RemoteAddr())
				reader := bufio.NewReader(conn)

				for {
					line,_,err := reader.ReadLine()
					if err != nil {
						log.Println(err)
						break
					}else {
						if string(line) == "quit"{
							break
						}
						log.Println("接收到数据",string(line))
						fmt.Fprintf(conn,"Over %s\n",time.Now().Format("2006-01-02 15:04:05"))
					}
				}

			}()
	}
}
