package main

import (
	"html/template"
	"os"
	"strings"
)

func main()  {
	tplText := `{{upper .}}`

	funcs := template.FuncMap{
		"upper": strings.ToUpper,
	}
	tpl :=template.Must(template.New("tpl").Funcs(funcs).Parse(tplText))
	tpl.Execute(os.Stdout,"abcdefg")
}
