package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func postReq(i int) {
	//----Request----------------------
	resp, err := http.PostForm("http://localhost:9090/",
		url.Values{"foo": {string(i)}, "id": {"123"}})
	if nil != err {
		fmt.Println("errorination happened getting the response", err)
		return
	}

	//---Reply-----------------------
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		fmt.Println("errorination happened reading the body", err)
		return
	}
	fmt.Println(string(body[:]))
}

func main() {
	for i := 0; i < 100; i++ {
		postReq(i)
	}
}
