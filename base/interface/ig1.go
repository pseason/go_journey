/*
go 简单http服务器
*/
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "go 简单http服务器")
}

func home(w http.ResponseWriter, r *http.Request) {
	content, _ := ioutil.ReadFile("./home.html")
	w.Write(content)
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello go")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/sayHello", sayHello)
	http.HandleFunc("/home", home)
	log.Fatal(http.ListenAndServe("localhost:8986", nil))
}
