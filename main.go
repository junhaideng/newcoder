package main

import (
	"fmt"

	"github.com/junhaideng/newcoder/scrapy"
)

func main() {
	n := scrapy.New(scrapy.WithQuery("百度"))
	err := n.SaveJson("data.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("save successfully")
}
