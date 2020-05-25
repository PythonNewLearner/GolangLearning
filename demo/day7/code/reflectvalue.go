package main

import (
	"fmt"
	"reflect"
	"time"
)

func main()  {
	var a int
	var es = []interface{}{1,"test",true, time.Now(),&a,func() {}}
	for _,e :=range es{
		val:=reflect.ValueOf(e)
		fmt.Println(val.Type(),val.CanSet())

	}
}
