package main

import (
	"fmt"
	"time"
)

func main()  {
	var interval time.Duration = time.Second*3
	for {
		fmt.Println(<-time.After(interval))
	}


}



