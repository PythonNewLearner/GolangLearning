package main

import (
	"fmt"
	"net"
)

func main()  {
	inters,_:=net.Interfaces()
	for _,inter := range inters {
		fmt.Println(inter.Name,inter.HardwareAddr,inter.MTU)
	}
}
