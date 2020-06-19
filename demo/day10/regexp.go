package main

import (
	"fmt"
	"regexp"
)

func main()  {
	reg,err := regexp.Compile("132\\d{8}")
	fmt.Println(reg,err)

	fmt.Println(reg.MatchString("132xxx"))
	fmt.Println(reg.ReplaceAllString("11113287654321","999999"))

}
