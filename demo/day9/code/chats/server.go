package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
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
				scanner := bufio.NewScanner(os.Stdin)

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
						fmt.Println("请输入数据:")
						scanner.Scan()
						fmt.Fprintf(conn,"Over %s\n",scanner.Text())
					}
				}

			}()
	}
}
