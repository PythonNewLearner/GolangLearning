package main

import (
	"fmt"
	"strconv"

)

func main()  {
	ascii := "abc前后端一起来"
	fmt.Println([]byte(ascii))   //字节切片 
	fmt.Println([]rune(ascii))   //字节切片   int32

	fmt.Println(len(ascii))
	fmt.Println(len([]rune(ascii)))     // length
	fmt.Println(string([]rune(ascii)))   // convert back to normal string from rune 
	fmt.Println(string([]byte(ascii)))

	//fmt.Println(utf8.RuneCountInString(ascii))

	fmt.Println(strconv.FormatFloat(3.1415926,'f',6,64))
	fmt.Println(strconv.ParseFloat("3.1415926",64))

	fmt.Println(strconv.FormatBool(true))
	fmt.Println(strconv.ParseBool("true"))

	fmt.Println(strconv.ParseInt("88",10,8))

	fmt.Println(strconv.FormatInt(88,2)) // to binary number in string format



}