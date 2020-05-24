package main

import "fmt"

func main()  {
	var (
	achar byte = 'A'
	aint byte = 65
	unicodePoint rune = '移'

	)

	fmt.Println(achar,aint,unicodePoint)

	fmt.Printf("%U",unicodePoint)

	
	
}