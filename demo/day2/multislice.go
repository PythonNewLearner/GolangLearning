package main
import (
	"fmt"
)

func main()  {
	multi := [][]int{}
	fmt.Printf("%T %#v\n",multi,multi)

	multi = append(multi,[]int{1,2,3,4})
	multi = append(multi, []int{1,2,3})
	fmt.Printf("%#v\n",multi)
}