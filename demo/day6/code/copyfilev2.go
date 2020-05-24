package main
import (
	"fmt"
	"os"
	"io"
)

func CopyFile(src ,dest string) error  {
	srcFile,err1 := os.Open(src)
	if err1 !=nil {
		return err1
	}
	defer srcFile.Close()

	destFile,err2 := os.Create(dest)
	if err2 != nil {
		return err2
	}
	defer destFile.Close()

	_,err := io.Copy(destFile,srcFile)

	// _,err := io.CopyN(destFile,srcFile,10)   拷贝前N个字节

	return err
}


func main()  {

	fmt.Println(CopyFile("os.go","os.go.3"))
	
}