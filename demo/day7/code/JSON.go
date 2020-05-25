package main

import (
	"encoding/json"
	"fmt"
	"time"
)

const LayOut = "2006-01-02 15:04:05"
type Birthday *time.Time

func (b *Birthday) UnmarshalText(txt []byte) error  {
	t,_ := time.Parse(LayOut,string(txt))
	*b = &t
}

type User struct {
	Id int  `json:"id"`
	Name string
	Password string  `json:"-"`
	Sex bool
	Birthday *time.Time
}

func main()  {
	now := time.Now()
	user := []User{
		{1,"a","123",true,&now},
		{2,"b","345",false,&now},
		{3,"c","789",true,&now},
	}
	fmt.Println(user)
	b,_ :=json.MarshalIndent(user,"","\t")
	fmt.Println(string(b))
}
