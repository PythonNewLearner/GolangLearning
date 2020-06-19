package main

import (
	"fmt"
	"html/template"
	"os"
)

func main()  {
	tplText := "{{ .Name }}"
	tpl := 	template.Must(template.New("tpl").Parse(tplText))


	tpl.Execute(os.Stdout, struct {
		Name string
		Sex string
	}{
		"baichen",
		"male",
	})
	fmt.Println()



}
