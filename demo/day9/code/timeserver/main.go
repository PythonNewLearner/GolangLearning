package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main()  {


	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		ctx,err := ioutil.ReadFile("index.html")
		if err == nil {
			response.Write(ctx)
		}else {
			fmt.Fprint(response,"welcome")
		}
	})
	http.ListenAndServe(":9888",nil)
}
