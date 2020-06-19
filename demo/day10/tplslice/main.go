package main

import (
	"fmt"
	"html/template"
	"os"
)

func main()  {
	tplText := "{{ index . 1 }}"
	tpl := 	template.Must(template.New("tpl").Parse(tplText))


	tpl.Execute(os.Stdout,[]int{1,2,3,4,5})
	fmt.Println()


}
