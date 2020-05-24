package main

import (
	"os"
	"log"
)

func main()  {
	logfile,err := os.OpenFile("append.log",os.O_CREATE|os.O_APPEND|os.O_WRONLY,os.ModePerm)
	if err != nil{
		return
	}
	defer logfile.Close()
	log.SetOutput(logfile)
	log.Println("test")

}