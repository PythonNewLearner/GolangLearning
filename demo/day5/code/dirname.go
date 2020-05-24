package main

import (
	"os"
	"fmt"
)

func main()  {
	dir,err := os.Open("homework")
	fmt.Println(dir)
	if err != nil  {
		return
	}
	defer dir.Close()

	names,err := dir.Readdirnames(-1)
	fmt.Println(names,err)
}