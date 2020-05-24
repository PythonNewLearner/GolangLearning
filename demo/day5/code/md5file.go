package main

import (
	"os"
	"flag"
	"io"
	"crypto/md5"
	"fmt"
)

func main()  {
	var name string
	flag.StringVar(&name,"p","","path")
	flag.Parse()

	if name == ""{
		return
	}
	file,err := os.Open(name)
	if err != nil {
		return
	}
	defer file.Close()
	hasher := md5.New()
	ctx := make([]byte,1024)

	for {
		n,err := file.Read(ctx)
		if err == io.EOF {
			break
		}
		hasher.Write(ctx[:n])
	}
	fmt.Printf("%x\n",hasher.Sum(nil))
}