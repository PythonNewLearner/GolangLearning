package main

import (
	"fmt"
	"bytes"
	"os"
)
func main()  {
	buffer1 := bytes.NewBuffer([]byte("OK gogogo,"))
	buffer2 := bytes.NewBufferString("GOGOGO!!!,")

	buffer1.Write([]byte("in Shanghai"))
	buffer2.Write([]byte("in Dalian"))

	buffer1.WriteString("Op1")
	buffer2.WriteString("Op2")

	ctx := make([]byte,3)
	n,_ := buffer1.Read(ctx)
	fmt.Println(n,string(ctx[:n]))

	n,_ = buffer2.Read(ctx)
	fmt.Println(n,string(ctx[:n]))

	ctx,_= buffer2.ReadBytes(',')  //读取到逗号
	fmt.Println(string(ctx))

	txt,_ := buffer1.ReadString(',')  ////读取到逗号
	fmt.Println(txt)

	fmt.Println(string(buffer1.Bytes()))  //把剩余的转为切片
	fmt.Println(buffer2.String())  //把剩余的转为字符串

	buffer1.WriteTo(os.Stdout)      //输出剩余的内容

	buffer2.Reset() //清空内容
	buffer2.WriteTo(os.Stdout)
}