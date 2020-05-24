package main
import (
	"fmt"
	
)

//panic 抛出错误
func test() (err error) {
	defer func ()  {
		fmt.Println("defer")
		//recover处理了错误panic
		if panic_err := recover();panic_err != nil{
			
			err = fmt.Errorf("%s",panic_err)
			
		}
	}()
	fmt.Println("before")
	panic("自定义panic")  
	fmt.Println("after")
	return
}
func main()  {
	test()

	
}