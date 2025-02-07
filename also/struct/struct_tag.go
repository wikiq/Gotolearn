package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	article := Article{
		Title:    "Golang for Qiws",
		Desc:     "desc123123",
		Content:  "content1231231",
		Author:   "author12312312",
		UserName: "admin123",
		password: "password123",
	}
	// json.Marshal(article)
	jsonData, err := json.Marshal(article)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonData))
}

type Article struct {
	Title    string
	Desc     string
	Content  string
	Author   string
	UserName string
	password string
}
