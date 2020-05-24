package main

import (
	"fmt"
	"io/ioutil"
)
func main()  {
	//创建一个临时文件，父目录为./test,文件命名为以log为前缀+随机数
	//返回文件夹路径,error
	dir,_ := ioutil.TempDir("./test","log")
	fmt.Println(dir)
}