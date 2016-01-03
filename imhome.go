package main

import (
    "fmt"
    "net/http"
    "strings"
)

var users = map[string]string{}

func handler(w http.ResponseWriter, r *http.Request) {
    ip := strings.Split(r.RemoteAddr, ":")[0]
    users[ip] = ip
    fmt.Fprint(w, users)
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("listing at http://localhost:8080/")
    http.ListenAndServe(":8080", nil)
}
