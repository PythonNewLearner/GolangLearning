package main

import (
	"fmt"
	"strconv"
	"errors"
)

func div(n1,n2 int) (int ,error)  {
	if n2 == 0{
		return -1, errors.New("除数不能为零")  //自定义错误 
	}
	return n1/n2,nil
}

func main()  {
	value,err := strconv.Atoi("xxx")
	fmt.Println(value)
	fmt.Printf("%#v\n",err)   // err is pointer type

	e := fmt.Errorf("自定义错误1")
	fmt.Printf("%T %#v\n",e,e)

	e = errors.New("自定义错误2")
	fmt.Printf("%T %#v\n",e,e)

	if rt, err := div(1,0);err == nil{
		fmt.Println(rt)
	} else{
		fmt.Println(err)
	}
}

