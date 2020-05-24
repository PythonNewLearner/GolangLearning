package main
import (
	"io/ioutil"
	"time"
)

func main()  {
	file,_ := ioutil.TempFile("./test","log")
	file.WriteString(time.Now().Format("15:04:05"))
	file.Close()
}