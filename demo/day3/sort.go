package main

import (
	"fmt"
	"sort"
)

func main()  {
	stats := [][]int{{'A',3},{'B',2},{'C',4},{'D',1}}
	sort.Slice(stats,func(i,j int) bool { return stats[i][1] > stats[j][1]})
	fmt.Println(stats)

	index := sort.Search(len(stats),func(i int) bool {return stats[i][1]<=1}) // ascending order use >=; descending order use <=
	fmt.Println(index)

}


