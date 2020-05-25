package main

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

func GetType( t reflect.Type) string  {
	//reflect.Type
	switch t.Kind() {
	case reflect.Int,reflect.String,reflect.Float64,reflect.Bool:
		return t.Name()
	case reflect.Array:
		return fmt.Sprintf("[%d]%s",t.Len(),GetType(t.Elem()))
	case reflect.Slice:
		return fmt.Sprintf("[]%s",GetType(t.Elem()))
	case reflect.Map:
		return fmt.Sprintf("map[%s]%s",GetType(t.Key()),GetType(t.Elem()))
	case reflect.Struct:
		var builder strings.Builder
		builder.WriteString(fmt.Sprintf(" %s struct\n",t.Name()))
		builder.WriteString(fmt.Sprintf("Fields %d\n",t.NumField()))
		for i:=0;i<t.NumField();i++{
			field := t.Field(i)
			builder.WriteString(fmt.Sprintf("\t\t%s %s %s\n",field.Name,GetType(field.Type),field.Tag))
		}
		builder.WriteString(fmt.Sprintf("Number of methods %d \n",t.NumMethod()))
		for i:=0;i<t.NumMethod();i++{
			method := t.Method(i)
			builder.WriteString(fmt.Sprintf(" Method : Name (%s)\tType (%s)\tFunc(%s)\n ",method.Name,method.Type,method.Func))
		}
		return builder.String()
	case reflect.Ptr:
		var builder strings.Builder
		builder.WriteString(fmt.Sprintf("Pointer %s",GetType(t.Elem())))
		return builder.String()
	default:
		return "Unknown"
	}
}

func main()  {
	var es = []interface{}{1,"test",false,[2]int{1,2},
		[]int{1,2,3},
		map[string]int{"id":1},
		time.Now(),
		3.12,
		[2][]int{{100,90,80},{100,90,80}},
		[][]int{{2,3,4},{1,5,6},{7,8,9}},
	}
	for _,e := range es{
		fmt.Println(GetType(reflect.TypeOf(e)))
	}
}
