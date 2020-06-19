package main

import (
	"html/template"
	"os"
)

func main()  {
	tpl := template.Must(template.ParseFiles("html/index.html","html/test.html"))
	tpl.ExecuteTemplate(os.Stdout,"index.html",[]int{1,2,3,4,5})
	//tpl.ExecuteTemplate(os.Stdout,"test.html",[]int{1,2,3,4,5})

}
