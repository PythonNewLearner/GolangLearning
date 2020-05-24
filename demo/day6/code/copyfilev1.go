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

	ctx := make([]byte,10)
	for {
		n,e := srcFile.Read(ctx)
		if e == io.EOF{
			break
		}
		n,e = destFile.Write(ctx[:n])
		if e != nil {
			return e
		}
	}

	return nil
}


func main()  {

	fmt.Println(CopyFile("os.go","os.go.2"))
	
}