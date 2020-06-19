package main

import "net/http"

func main()  {
	addr := ":9888"
	http.Handle("/static/",http.StripPrefix("/static/",http.FileServer(http.Dir("www"))))
	http.ListenAndServe(addr,nil)
}
