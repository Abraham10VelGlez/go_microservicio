package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)


func main() {

    http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello,", html.EscapeString(r.URL.Path))
    })

    http.HandleFunc("/mensaje", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Go service, para Abraham10VelGlez")
    })


	http.HandleFunc("/metodoget", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, " creando api go")
    })

    log.Fatal(http.ListenAndServe(":8081", nil))
}