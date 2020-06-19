package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
)

func main()  {
	url := "https://godoc.org/?q=" + "query"

	document , err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}
	selection := document.Find(".table-condensed").Find("a") // by class
	selection.Each(func(index int,tag *goquery.Selection){
		href,exists :=tag.Attr("href")

		fmt.Println(index,tag.Text(),href,exists)
	})

	selection = document.Find("#x-search")  // by id
	fmt.Println(selection.Attr("class"))
	fmt.Println(selection.Text())
	fmt.Println(selection.Html())

	selection = document.Find(".table-condensed a")  //子孙选择器
	selection.Each(func(index int,tag *goquery.Selection){
		href,exists :=tag.Attr("href")

		fmt.Println(index,tag.Text(),href,exists)
	})
}
