package main

import "fmt"
//空接口
type EmptyIntf interface {
}

func PrintType(v interface{})  {
	switch value :=v.(type) {
	case int:
		fmt.Println("int",value)
	case string:
		fmt.Println("string",value)
	case bool:
		fmt.Println("Bool",value)
	case []int:
		fmt.Println("Slice int",value)
	case *int:
		fmt.Println("Pointer int",value)
	case map[string]int:
		fmt.Println("Map[string]int ",value)
	default:
		fmt.Println("Unknown type")
	}
}
func main()  {
	var empty EmptyIntf
	empty = 1
	empty = true
	empty = "test"
	fmt.Printf("%#v\n",empty)

	var empty1 interface{}
	empty1 = 1
	empty1 = false
	empty1 = "test"
	fmt.Println(empty1)

	PrintType([]int{1,2,3,4})
	PrintType(23)
	PrintType("USA")
	PrintType(map[string]int{"USA":1})

}
