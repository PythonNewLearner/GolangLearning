package main

import (
	"fmt"
	"os"
	"strings"
)
type FileFilter func(string) bool

type Callback func(string)

func Dir(path string,filter FileFilter,callback Callback)  {
	file,err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	fileinfos, err := file.Readdir(-1)
	if err != nil {
		return
	}
	for _, fileinfo := range fileinfos{
		fpath := path + "/" + fileinfo.Name()
		
			if fileinfo.IsDir() {
				Dir(fpath,filter,callback)
			}
			if filter == nil || filter (fpath){
				if callback != nil{
					callback(fpath)
				}
			}	
	}
}
func main()  {	
	Dir("homework",func (path string) bool {
		return strings.HasSuffix(path,".go")
	},func (path string)  {
		fmt.Println(path) //???
	})

}