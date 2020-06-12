package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

func sayHello(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
		fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username"))) //サーバ側に出力されます。
		fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
		template.HTMLEscape(w, []byte(r.Form.Get("username")))
	}
}

func main(){
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/login", login)
	server.ListenAndServe()
}