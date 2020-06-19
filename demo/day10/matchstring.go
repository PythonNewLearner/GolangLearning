package main

import (
	"fmt"
	"regexp"
)

func main()  {
	pattern := "^132\\d{8}$"
	fmt.Println(regexp.MatchString(pattern,"13245654"))
	fmt.Println(regexp.MatchString(pattern,"13245xxx"))
	fmt.Println(regexp.MatchString(pattern,"132123456780"))

	pattern = "^(132)|(158)][0-9]{8}$"
	fmt.Println(regexp.MatchString(pattern,"132123456780"))

	pattern = "[a-zA-Z0-9]{1,32}@[a-z]{1,12}[.]edu$"
	fmt.Println(regexp.MatchString(pattern,"baichen@qq.edu"))

	fmt.Println(regexp.QuoteMeta(pattern))

}
