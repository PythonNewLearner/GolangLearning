package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main()  {
	conn,err :=net.Dial("tcp","127.0.0.1:9999")
	if err!= nil{
		log.Fatal(err)
	}
	defer conn.Close()
	log.Println("connected")
	scanner := bufio.NewScanner(os.Stdin)
	reader := bufio.NewReader(conn)
	for i:=0;i<4;i++{
		fmt.Println("请输入内容:")
		scanner.Scan()

		fmt.Fprintf(conn,"%s\n",scanner.Text())
		line,_,err :=reader.ReadLine()
		if err != nil {
			log.Println(err)
			break
		}
		log.Println(string(line))
	}


	fmt.Fprintln(conn,"quit")





}