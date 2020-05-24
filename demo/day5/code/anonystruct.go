package main
import (
	"fmt"
)
func main()  {
	//匿名结构体 : 直接初始化给一个变量
	var user struct{  //这是一个实例了
		id int
		name string
		age int
	}
	fmt.Printf("%T\n",user)
	fmt.Printf("%#v\n",user)
	fmt.Println(user.name)
	user.name = "chen"
	fmt.Printf("%#v\n",user)

	user = struct{
		id int
		name string
		age int 
	}{1,"chen",30}
	fmt.Printf("%#v\n",user)



	

}