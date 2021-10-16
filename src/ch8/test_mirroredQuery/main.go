package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func mirroredQuery() (str1 string, str2 string, str3 string) {
	responses := make(chan string, 3)
	go func() { responses <- request("https://www.gd.gov.cn/") }()
	go func() { responses <- request("https://www.zj.gov.cn/") }()
	go func() { responses <- request("https://www.ah.gov.cn/") }()
	return <-responses, <-responses, <-responses
}

func request(hostname string) (response string) {
	resp, err := http.Get(hostname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", hostname, err)
		os.Exit(1)
	}
	response = string(b)
	return response
}

func main() {
	str1, str2, str3 := mirroredQuery()
	fmt.Println(str1)
	fmt.Println(str2)
	fmt.Println(str3)
}
