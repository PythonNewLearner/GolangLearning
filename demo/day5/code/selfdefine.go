package main
import(
	"fmt"
)

type Counter int
type Task map[string]string
type Callback func(...string)
func main()  {
	var cnt Counter
	fmt.Printf("%T\n",cnt)
	fmt.Printf("%#v\n",cnt)
	cnt = 1
	fmt.Println(cnt)
	
	// var total int = 10
	// fmt.Println(cnt/total)     // mismatched type

	var task Task
	fmt.Printf("%#v\n",task)
	task = map[string]string{"name":"Good task"}
	fmt.Printf("%#v\n",task)

	var call Callback
	call = func(args ... string)  {
		for _,v := range args{
			fmt.Println(v)
		}
	}
	call("a","b","c")




}