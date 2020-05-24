package main

import (
	"fmt"
)

func main()  {
	var scores map[string]float64   // nil
	//scores["pp"] = 87 error as it is a nil map 
	fmt.Printf("%T , %#v\n",scores,scores)

	scores = map[string]float64{}  // empty
	fmt.Printf("%T , %#v\n",scores,scores)
	scores["pp"] = 87   // no error as it is empty map
	fmt.Printf("%T , %#v\n",scores,scores)

	scores = map[string]float64{"1":90,"3":43,"5":80}
	fmt.Printf("%T , %#v\n",scores,scores)

	//make
	scores = make(map[string]float64)   // empty
	fmt.Printf("%T , %#v\n",scores,scores)

	scores = map[string]float64{"1":90,"3":43,"5":80}
	fmt.Println(len(scores))


	fmt.Println(scores["1"])

	fmt.Println(scores["xxx"])    // return 0 value

	// if element in the map
	k,ok := scores["yy"]
	fmt.Println(k,ok)   // ok =false: not exist in the map

	scores["xx"] = 100
	fmt.Println(scores)
	scores["1"] = 99
	fmt.Println(scores)

	delete(scores,"xx")
	fmt.Println(scores)
	delete(scores,"bb")
	fmt.Println(scores)

	for k,v := range scores{
		fmt.Println(k,v)
	}

}