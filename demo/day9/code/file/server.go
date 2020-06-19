package main

import (
	"log"
	"net"
	"os"
)

func handleConn(conn net.Conn)  {
	defer conn.Close()
}

func main()  {
	logfile,_ := os.OpenFile("fileserver.log",os.O_CREATE|os.O_APPEND|os.O_RDWR,os.ModePerm)
	defer logfile.Close()

	log.SetOutput(logfile)

	addr := "0.0.0.0:9999"  // 监听所有网卡的9999
	listener,err:= net.Listen("tcp",addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	for {
		conn,err :=listener.Accept()
		if err != nil{
			log.Printf("%s\n",err)
			continue
		}
		log.Printf("Client connected : %s\n",conn.RemoteAddr())
	}
}
