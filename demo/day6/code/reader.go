package main

import (
	"fmt"
	"strings"
	"io"
	"os"
)

func main()  {
	reader:=strings.NewReader("12345abcdefg")

	ctx := make([]byte,3)
	for {
		n,err := reader.Read(ctx)
		if err == io.EOF{
			break
		}
		fmt.Println(n,err,string(ctx))
	}
	reader.Seek(0,os.SEEK_SET)
	n,err := reader.Read(ctx)
	fmt.Println(n,err,string(ctx[:n]))

	fmt.Println(reader.Size())
	reader.Reset("3123123")

	n,err = reader.Read(ctx)
	fmt.Println(n,err,string(ctx[:n]))

	reader.WriteTo(os.Stdout)



}