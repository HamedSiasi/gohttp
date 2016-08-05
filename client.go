package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	//----Request----------------------
	url := "http://localhost:9090/"
	fmt.Println("URL:>", url)

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	//---Reply-----------------------
	fmt.Println("Status:", resp.Status)
	fmt.Println("Headers:", resp.Header)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Body:", string(body))
}
