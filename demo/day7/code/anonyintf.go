package main

import "fmt"

type EmailSender struct {

}

func (s EmailSender)Send(to string,msg string) error  {
	return nil
}
func main()  {
	var sender interface{
		Send(string,string)error
	}
	fmt.Printf("%T %#v\n",sender,sender)
	sender = EmailSender{}
	fmt.Printf("%T %#v\n",sender,sender)
}









