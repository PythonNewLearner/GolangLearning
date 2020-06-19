package main

import (
	"fmt"
	"net"
)

func main()  {
	fmt.Println(net.JoinHostPort("123.0.0.1","6666"))
	fmt.Println(net.JoinHostPort("::1","6666"))
	host,port,err :=net.SplitHostPort("127.0.0.1:9999")
	fmt.Println(host,port,err)

	fmt.Println(net.LookupAddr("127.0.0.1"))  //通过IP找主机名
	fmt.Println(net.LookupHost("localhost"))  //通过主机名找IP

	//IP
	ip := net.ParseIP("127.0.0.1")
	fmt.Printf("%#v\n",ip)

	ip = net.ParseIP("::1")
	fmt.Printf("%#v\n",ip)

	ipaddr,e := net.LookupIP("www.google.com")
	fmt.Println(ipaddr,e)

	ip,ipnet,err := net.ParseCIDR("192.168.1.1/24")
	fmt.Println(ip)
	fmt.Println(ipnet)
	fmt.Println(err)
	fmt.Println(ipnet.Contains(ip))
	fmt.Println(ipnet.Contains(net.ParseIP("192.168.2.1")))
}
