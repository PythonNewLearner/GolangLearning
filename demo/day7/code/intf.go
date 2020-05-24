package main

import "fmt"

type Sender interface {
	Send(string,string) error
	SendAll([]string,string) error
}
type EmailSender struct {

}

func (s EmailSender)Send(to , msg string) error  {
	fmt.Println("To:",to,"Message:",msg)
	return nil
}
func (s EmailSender)SendAll(tos []string,msg string) error  {
	fmt.Println("To those: ",tos,"Message: ",msg)
	return nil
}
func main()  {
	var sender Sender
	fmt.Printf("%T %#v\n",sender,sender)

	sender = EmailSender{}
	fmt.Printf("%T %#v\n",sender,sender)
	sender.Send("Chen","Good morning!")
	sender.SendAll([]string{"chen","Trump","BoMing"},"Hello US")

}

//注：接口不能访问属性