package main
import (
	"fmt"
	"path/filepath"
	"os"
)
func main()  {
	filepath.Walk(".",func(path string,file os.FileInfo,err error) error  {
		fmt.Println(path,file.Name())
		return nil
	})
}