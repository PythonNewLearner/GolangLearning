package main
import (
	"fmt"
	"strings"

)

func main()  {
	fmt.Println(strings.Compare("a","b"))  // "a"<"b"

	fmt.Println(strings.Contains("thisGo","Go"))  // true 

	fmt.Println(strings.ContainsRune("前后端一起来",'来'))  //true


	fmt.Println(strings.Count("前后端一起来","来"))   // 1

	fmt.Println(strings.EqualFold("abc","ABC"))   //true

	fmt.Printf("%#v\n",strings.Fields("a b\tc\nd\re\ff"))  // remove spaces , return slice:   []string{"a", "b", "c", "d", "e", "f"}

	fmt.Println(strings.HasPrefix("abc","ab"))  //start with ab: return true

	fmt.Println(strings.HasSuffix("abcd","cd"))  //end with cd: true


	fmt.Println(strings.Index("abcdefg","bc"))   // return index of bc in the string
	fmt.Println(strings.Index("abcdefg","xx"))  // return -1 if xx not in the string

	// Join
	fmt.Println(strings.Join([]string{"ab","cd","ef"},"-"))  // join strings with "-"

	// Split
	fmt.Printf("%#v\n",strings.Split("ab,cd,ef",","))  // split with  ","
	fmt.Printf("%#v\n",strings.SplitN("ab,ab,ab,ab",",",2))  // split first 2


	//repeat
	fmt.Println(strings.Repeat("*",5))

	//replace
	fmt.Println(strings.Replace("xyxyfdfdzzkk","kk","mn",-1))   // replace all kk with mn
	fmt.Println(strings.Replace("xyxyfdfdzzkk","yx","mn",1))    // replace 1st yx with mn
	fmt.Println(strings.ReplaceAll("xyxyfdfdzzkk","xy","mn"))   //replace all xy with mn


	//trim
	fmt.Println(strings.Trim("abcdefgabc","abc"))  // trim "abc" in the beginning and end
	fmt.Println(strings.TrimSpace(" \t\n\fabcdef\t\n"))   // trim spaces
	fmt.Println(strings.TrimLeft("abcdefgabc","abc"))    // trim left
	fmt.Println(strings.TrimRight("abcdefgabc","abc"))    // trim right

}