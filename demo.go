package main

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

func main() {
	url := "http://api.tiankongapi.com/api.php/provide/vod?ac=list&t=6&pg=1"

	args := &fasthttp.Args{}
	// args.Add("ac", "list")
	// args.Add("t", "6")
	// args.Add("pg", "1")

	statusCode, respBody, err := fasthttp.Post(nil, url, args)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	fmt.Printf("Status Code: %d", statusCode)
	fmt.Printf("Response Body: %s", respBody)
}
