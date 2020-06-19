package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func ls()  {

}
func main()  {
	logfile,_ := os.OpenFile("client.log",os.O_APPEND|os.O_CREATE|os.O_WRONLY,os.ModePerm)
	defer logfile.Close()
	log.SetOutput(logfile)

	addr := "127.0.0.1:9999"
	conn,err :=net.Dial("tcp",addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	log.Println("connected to fileserver")
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("请输入指令:")
		scanner.Scan()
		input := scanner.Text()
		cmds := strings.Split(input," ")
		switch cmds[0] {
		case "ls":
			
		
		}
	}

}
