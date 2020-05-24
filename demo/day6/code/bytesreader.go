package main

import (
	"bytes"
	"fmt"
	"os"
)

func main()  {
	reader := bytes.NewReader([]byte("abcdef"))
	ctx := make([]byte,5)
	n,err := reader.Read(ctx)
	fmt.Println(n,err,string(ctx))
	reader.WriteTo(os.Stdout)  //输出剩余的内容
}