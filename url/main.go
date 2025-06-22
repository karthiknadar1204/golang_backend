package main

import (
	"fmt"
	"net/url"
)

func main() {
	urlString := "https://www.google.com/search?q=golang"
	url,err := url.Parse(urlString)
	if err != nil {
		fmt.Println("error",err)
		return
	}
	fmt.Println(url.Scheme)
	fmt.Println(url.Host)
	fmt.Println(url.Path)
	fmt.Println(url.Query())
}
