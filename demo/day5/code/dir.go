package main
import (
	"os"
)

func main()  {
	os.Mkdir("test",os.ModePerm)
	os.MkdirAll("test/a/b",os.ModePerm)
	os.Remove("test")
	os.RemoveAll("test")
}