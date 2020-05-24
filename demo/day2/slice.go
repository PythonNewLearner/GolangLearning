package main
import (
	"fmt"
)

func main()  {
	var names []string   // nil value
	fmt.Printf("%T\n",names)
	fmt.Printf("%#v\n",names)

	names = []string{}  // empty
	fmt.Printf("%#v\n",names)

	names = []string{"a","b","c"}
	fmt.Printf("%#v\n",names)

	names = []string{1:"a",100:"kk"}
	fmt.Printf("%#v\n",names)


	// make slice
	names = make([]string,5)  //  make a slice with 5-elements, cap is 5 if no cap value passed
	fmt.Printf("%#v\n",names)

	names = make([]string,0,10) // 0 elements, 10 capacity
	fmt.Printf("%#v\n",names)

	names = make([]string,3,10)//  3 length, 10 cap before asking for new memory space

	fmt.Println(len(names),cap(names))  //  len:3 , cap:10

	// add elements
	names = append(names,"d")
	fmt.Printf("%#v\n",names)

	for i,v := range names{
		fmt.Println(i,v)
	}

	//copy
	aSlice := []string{"a","b","c"}
	bSlice := []string{"x","y"}

	copy(aSlice,bSlice)
	fmt.Println(aSlice,bSlice)  // copy bSlice to aSlice

	nums := make([]int,6,10)
	numsChild := nums[3:4]
	fmt.Println(numsChild)  // end <=cap

	numsChild = append(numsChild,100)   // when numChild changes, nums also changes
	fmt.Println(numsChild)
	fmt.Println(nums)

	// nums[start:end:max]   start<=end<=max<=cap, new_cap=max-start
	nums = []int{0,1,2,3,4,5,6}
	numsChild = nums[3:4:4]      // new_cap=max-start, newcap = 4-3 =1 => will get a new space when append
	fmt.Println(cap(numsChild))
	numsChild = append(numsChild,100)
	fmt.Println(cap(numsChild))
	fmt.Println(nums,numsChild)

	//array
	numsArray := [6]int{0,1,2,3,4,5}
	arrayChild := numsArray[3:4] // start<=end<=len
	fmt.Printf("%T %#v\n",arrayChild,arrayChild)
	fmt.Println(cap(arrayChild))    // new_cap = len -start = 6 - 3 =3

	arrayChild = append(arrayChild,100)
	fmt.Println(numsArray,arrayChild)

	numsArray = [6]int{0,1,2,3,4,5}
	arrayChild = numsArray[3:4:4]    //new_cap=max-start, newcap = 4-3 =1 => will get a new space when append
	arrayChild = append(arrayChild,200)    // no effect on original array
	fmt.Println(numsArray,arrayChild)

	










}





