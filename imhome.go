package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello Go Webserver")
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("listing at http://localhost:8080/")
    http.ListenAndServe(":8080", nil)
}
