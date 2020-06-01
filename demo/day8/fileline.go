package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func fileline(path string) int {
	count:=0
	file,err := os.Open(path)
	if err !=nil {
		return count
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line,_,err :=reader.ReadLine()
		if err == io.EOF{
			break
		}
		ctx := strings.TrimSpace(string(line))
		if ctx == "" || strings.HasPrefix(ctx,"//"){
			continue
		}
		count ++
	}
	return count
}
//计算每个go 文件的行数
func main()  {
	dir := "./../../"
	total := 0
	var wg sync.WaitGroup

	//walk遍历文件夹
	filepath.Walk(dir,func(path string, info os.FileInfo, err error) error{
		if filepath.Ext(path) == ".go" && !info.IsDir(){
			wg.Add(1)
			go func() {
				total +=  fileline(path)
				wg.Done()
			}()
		}
		return nil
	})
	wg.Wait()
	fmt.Println(total)
	fmt.Println(fileline("./account.go"))
}
