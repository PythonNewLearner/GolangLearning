package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main()  {
	conn,err :=net.Dial("tcp","127.0.0.1:9999")
	if err!= nil{
		log.Fatal(err)
	}
	defer conn.Close()
	log.Println("connected")
	reader := bufio.NewReader(conn)
	for i:=0;i<4;i++{

		fmt.Fprintf(conn,"Time %s\n",time.Now().Format("2006-01-02 15:04:05"))
		line,_,err :=reader.ReadLine()
		if err != nil {
			log.Println(err)
			break
		}
		log.Println(string(line))
	}


	fmt.Fprintln(conn,"quit")





}