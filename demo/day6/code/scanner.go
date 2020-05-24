package main

import (
	"os"
	"fmt"
	"bufio"
)
func main()  {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Input :")
	scanner.Scan()
	fmt.Println(scanner.Text())
}
