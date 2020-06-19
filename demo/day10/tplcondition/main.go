package main

import (
	"fmt"
	"html/template"
	"os"
)

func main()  {
	tplText := `
		{{ if eq .Sex 1 }}男{{ else }}女{{ end }}
			`
	tpl := 	template.Must(template.New("tpl").Parse(tplText))


	tpl.Execute(os.Stdout, struct {
		Name string
		Sex int

	}{
		"baichen",
		1,
	})
	fmt.Println()



}
