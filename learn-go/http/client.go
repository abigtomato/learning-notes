package main

import (
	"fmt"
	"regexp"
	"net/http"
	"net/http/httputil"
)

func main() {
	req, err := http.NewRequest(http.MethodGet, "http://www.zhenai.com/zhenghun/aba", nil)
	if err != nil {
		fmt.Printf("http.NewRequest() fail error: %v\n", err.Error())
		return
	}
	req.Header.Add("User-Agent", "")

	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("Redirect: %v\n", req)
			return nil
		},
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("http.DefaultClient.Do() fail error: %v\n", err.Error())
		return
	}
	defer resp.Body.Close()

	s, err := httputil.DumpResponse(resp, true)
	if err != nil {
		fmt.Printf("httputil.DumpResponse() fail error: %v\n", err.Error())
		return
	}

	fmt.Printf("content: %s\n", s)
}