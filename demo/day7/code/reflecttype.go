package main

import (
	"fmt"
	"reflect"
	"time"
)

func main()  {
	var es = []interface{}{1,"test",true, time.Now(),func() {}}
	for _,e :=range es{
		typ:=reflect.TypeOf(e)
		//fmt.Printf("%#v\n",typ)
		fmt.Println(typ.Name(),typ.PkgPath(),int(typ.Kind()),typ.NumMethod())

	}
}
