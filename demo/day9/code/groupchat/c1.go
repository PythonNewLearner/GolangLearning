package main

import (
	"bufio"
	"log"
	"net"
	"os"
)

func main()  {
	addr := "127.0.0.1:9888"
	for {
	conn,err := net.Dial("tcp",addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		conn.Write([]byte(scanner.Text()))
	}

}
