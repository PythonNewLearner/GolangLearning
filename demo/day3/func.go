package main
import (
	"fmt"
)

func add(a,b int) int {  // a and b are same type
	return a+b
}

func test(p1,p2 string,p3,p4 int,p5,p6 bool){
	fmt.Println(p1,p2,p3,p4,p5,p6)
}

func test1(args ... string)  {       // pass multiple args: args is slice type
	fmt.Printf("%T, %#v\n",args,args)
}

func test2(n1,n2 int,args ...int)  int {    // at least 2 : n1 n2; args ... must be at last
	total := n1 + n2
	for _,v := range args{
		total += v
	}
	return total
}

func calc(n1,n2 int,args ...int) int {
	return test2(n1,n2,args ...)   // Here args... is unpack
	
}

func main()  {
	fmt.Println(add(1,2))
	test("abc","efg",1,2,true,false)
	test1("ab","ef")  //args : []string{"ab", "ef"}
	fmt.Println(test2(1,2,3,4,5,6))
	fmt.Println(calc(1,2,3,4,5,6,7,8,9))
}