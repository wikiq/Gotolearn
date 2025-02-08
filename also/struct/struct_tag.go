package main

import (
	"fmt"
)

// func main() {
// 	article := Article{
// 		Title:    "Golang for Qiws",
// 		Desc:     "desc123123",
// 		Content:  "content1231231",
// 		Author:   "author12312312",
// 		UserName: "admin123",
// 		PassWord: "password123",
// 	}
// 	// json.Marshal(article)
// 	jsonData, err := json.Marshal(article)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(string(jsonData))
// }

// type Article struct {
// 	Title    string `json:"title"`
// 	Desc     string `json:"desc"`
// 	Content  string `json:"content"`
// 	Author   string `json:"author"`
// 	UserName string `json:"userName"`
// 	PassWord string `json:"password"`
// }

type error interface {
	Error() string
}

type RPCError struct {
	Code    int64
	Message string
}

func (e *RPCError) Error() string {
	return fmt.Sprintf("code: %d, message: %s", e.Code, e.Message)
}

func main() {
	var rpcErr error = NewRPCError(400, "unknown err") // typecheck1
	err := AsErr(rpcErr)                               // typecheck2
	println(err)
}

func NewRPCError(code int64, msg string) error {
	return &RPCError{ // typecheck3
		Code:    code,
		Message: msg,
	}
}

func AsErr(err error) error {
	return err
}
