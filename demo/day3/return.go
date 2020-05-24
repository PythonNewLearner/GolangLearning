package main
import (
	"fmt"
)

func calc(n1,n2 int) (int,int)  {
	return n1+n2,n1*n2
}

func calc1(n1,n2 int)(r1 int,r2 int)  {
	r1 = n1 + n2
	r2 = n1 * n2
	return r1,r2
}

func calc2(n1,n2 int)(r1 ,r2 int)  {
	r1 = n1 + n2
	r2 = n1 * n2
	return r1,r2
}

func main()  {
	fmt.Println(calc(3,4))
	fmt.Println(calc1(4,5))
	fmt.Println(calc2(4,5))
}



