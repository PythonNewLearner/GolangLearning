package main

import "fmt"

type SingleSender interface {
	Send(string , string) error
}

type Sender interface {
	SingleSender               //通过匿名嵌入进行扩展
	SendAll([]string,string) error
}
type EmailSender struct {

}
func (s EmailSender)SendAll(tos []string,msg string) error  {
	return nil
}
func (s EmailSender)Send(to string,msg string)error  {
	return nil
}
func main()  {
	var sender Sender = EmailSender{}
	fmt.Printf("%T %#v\n",sender,sender)
}
