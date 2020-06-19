package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main()  {
	conn,err :=net.Dial("udp","127.0.0.1:9998")
	if err != nil{
		log.Fatal(err)
	}
	defer conn.Close()
	fmt.Fprintf(conn,"%s\n",time.Now().Format("2006-01-02 15:04:05"))
	ctx := make([]byte,1024)
	for {
		n,_ :=conn.Read(ctx)
		fmt.Println(string(ctx[:n]))
	}

}
