package main

import (
	"fmt"
	"log"
	"net"
)

func main()  {
	addr := "0.0.0.0:9998"
	packetConn,err :=net.ListenPacket("udp",addr)
	if err != nil{
		log.Fatal(err)
	}
	defer packetConn.Close()

	ctx := make([]byte,1024)
	for {
		n,addr,err :=packetConn.ReadFrom(ctx)
		if err != nil{
			log.Println(err)
			continue
		}
		log.Println(addr)
		fmt.Println(string(ctx[:n]))
		packetConn.WriteTo([]byte("xxxxx"),addr)
	}
}
