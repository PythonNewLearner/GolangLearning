package main

import (

	"fmt"
)

type Sender interface {
	Send(string,string) error
	SendAll([]string,string) error
}
type EmailSender struct {
	name string
	addr string
	port int

}

func (s EmailSender)Send(to , msg string) error  {
	fmt.Println("To:",to,"Message:",msg)
	return nil
}
func (s EmailSender)SendAll(tos []string,msg string) error  {
	fmt.Println("To those: ",tos,"Message: ",msg)
	return nil
}

type SMSSender struct {
	id int
	key int
}

func (s SMSSender)Send(to , msg string) error {
	fmt.Println("SMS to :",to,"Message: ",msg)
	return nil
}
func (s SMSSender)SendAll(tos []string,msg string) error  {
	fmt.Println("SMS to all : ",tos,"Message: ", msg)
	return nil
}
func PrintConfig(sender Sender)  {
	switch v:= sender.(type) {
	case EmailSender:
		fmt.Println("EmailSender",v.name)
	case SMSSender:
		fmt.Println("SMSSender",v.id,v.key)
	case *SMSSender:
		fmt.Println("Pointer SMSSender ",v.id,v.key)
	default:
		fmt.Println("Wrong Type")
	}
}
func main()  {
	var sender Sender
	emailsender := EmailSender{"qq","深圳",123}
	sender = emailsender
	fmt.Printf("%T %#v\n",sender,sender)

	obj1,ok1 := sender.(EmailSender)  //断言
	fmt.Printf("%T %#v\n",obj1,obj1)
	fmt.Println(ok1)
	fmt.Println(obj1.addr,obj1.name,obj1.port)

	PrintConfig(sender)

	sender = SMSSender{id:1,key: 456}
	obj2,ok2 := sender.(EmailSender)   //断言
	fmt.Printf("%T %#v\n",obj2,obj2)
	fmt.Println(ok2)  //false  SMSSender 转不了EmailSender

	PrintConfig(sender)

	sender = &SMSSender{id:1,key: 456}
	obj3,ok3 := sender.(*SMSSender)    //断言
	fmt.Printf("%T %#v\n",obj3,obj3)
	fmt.Println(ok3)  //false  SMSSender 转不了EmailSender

	PrintConfig(sender)

}

