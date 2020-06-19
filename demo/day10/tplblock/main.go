package main

import (
	"html/template"
	"os"
)

func main()  {
	tplText := `{{block "content" . }} {{ . }} {{end}}`   //定义block名字叫做content
	tpl := template.Must(template.New("tpl").Parse(tplText))

	tpl,_ = tpl.Parse(`{{define "content"}}{{len .}}{{end}}`)  //覆盖block content
	tpl.Execute(os.Stdout,"abc")
}
