package main

import (
	"fmt"
	"path/filepath"
)
func main()  {
	path,_ := filepath.Abs(".")
	fmt.Println(path)  //   /home/chen/goproject/demo1/day6/code 
	fmt.Println(filepath.Base(path))  //code

	fmt.Println(filepath.Dir(path))   //parent dir
	fmt.Println(filepath.Clean("/../...////abc/adf"))

	path,_ = filepath.Abs("./filepath.go")
	fmt.Println(filepath.Ext(path))

	path1:= filepath.Dir(path)
	fmt.Println(filepath.HasPrefix(path,path1))

	dir,_ := filepath.Abs("/opt/cmdb")
	fmt.Println(filepath.Join(dir,"etc","app.ini"))

	fmt.Println(fmt.Println(filepath.Split(path)))

	fmt.Println(filepath.Glob("./test/a*"))
	fmt.Println(filepath.Glob("./test/*.txt"))



}