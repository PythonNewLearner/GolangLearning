package main

import (
	"fmt"
	"html/template"
	"os"
)

func main()  {
	tplText := "{{ .a }}"
	tpl := 	template.Must(template.New("tpl").Parse(tplText))


	tpl.Execute(os.Stdout,map[string]int{"a":1,"b":2,"c":3})
	fmt.Println()




}
