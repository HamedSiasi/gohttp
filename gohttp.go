package gohttp

import (
	"encoding/gob"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
	//"time"
	"io/ioutil"
)

// build a simple web server
// No need to Nginx or Apache
func Init() {

}

func RequestHandler1(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "count")
}

func RequestHandler2(res http.ResponseWriter, req *http.Request) {
	// parse and read http request headers and boy
	req.ParseForm()
	fmt.Println(req.Form)
	fmt.Println("path", req.URL.Path)
	fmt.Println("scheme", req.URL.Scheme)
	fmt.Println(req.Form["url_long"])
	for k, v := range req.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	// generate and write response
	fmt.Fprintf(res, "Hello")
}

func HttpServer() {
	http.HandleFunc("/", RequestHandler1)      // else
	http.HandleFunc("/count", RequestHandler2) //if path:/count
	// ...
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func HttpClient() {
	//----Request----------------------
	url := "http://localhost:9090/count"
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

func TcpServer() {
	//fmt.Println("tcpServer")
	// Create a TCP server
	// listen on a port
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		// accept a connection
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		// handle the connection
		go HandleServerConnection(c)
	}
}

func HandleServerConnection(c net.Conn) {
	//fmt.Println("handleServerConnection")
	// Receive the message
	var msg string
	err := gob.NewDecoder(c).Decode(&msg)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Received:", msg)
	}
	c.Close()
}

func TcpClient() {
	//fmt.Println("tcpClient")
	// connect to the server
	c, err := net.Dial("tcp", "localhost:9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	// send the message
	msg := "Hello World"
	fmt.Println("Sending:", msg)
	err = gob.NewEncoder(c).Encode(msg)
	if err != nil {
		fmt.Println(err)
	}
	c.Close()
}

//func main() {
// --- http server-client ---
//go httpServer()
//go httpClient()

// --- tcp server-client ---
//go tcpServer()
//go tcpClient()

//var input string
//fmt.Scanln(&input)
//}
