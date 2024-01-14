package main

import (
	_ "github.com/go-sql-driver/mysql"
)

type Article struct {
	Id          uint16
	Title       string `json:"name"`
	Anons       string `json:"anons"`
	Content     string `json:"content"`
	CategoryId  uint16 `json:"categoryId"`
	CreatedDate string `json:"createdDate"`
}

var posts []Article
var singlePost = Article{}

func main() {
	handleFunc()
}
