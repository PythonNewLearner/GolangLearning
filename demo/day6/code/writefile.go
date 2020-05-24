package main
import (
	"os"
	"io/ioutil"
)

func main()  {
	ioutil.WriteFile("test/name.txt",[]byte("chen"),os.ModePerm)
}