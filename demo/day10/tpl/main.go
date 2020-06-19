package main

import (
	"fmt"
	"html/template"
	"os"
)

func main()  {
	tplText := "{{ . }}"
	tpl := 	template.Must(template.New("tpl").Parse(tplText))
	tpl.Execute(os.Stdout,"baichen")
	fmt.Println()

	tpl.Execute(os.Stdout,[]int{1,2,3,4,5})
	fmt.Println()

	tpl.Execute(os.Stdout,map[string]int{"a":1,"b":2,"c":3})
	fmt.Println()

	tpl.Execute(os.Stdout, struct {
		Name string
		Sex string
	}{
		"baichen",
		"male",
	})
	fmt.Println()



}
