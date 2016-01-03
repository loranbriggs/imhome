package main

import (
    "fmt"
    "html/template"
    "net/http"
    "strings"
)

var users = map[string]string{}

func handler(w http.ResponseWriter, r *http.Request) {
    ip := strings.Split(r.RemoteAddr, ":")[0]
    users[ip] = ip

    data := struct {
        User string
        Users map[string]string
    } {
        ip,
        users,
    }

    t, _ := template.ParseFiles("views/home.html")
    t.Execute(w, data)
}

func newHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/new", newHandler)
    fmt.Println("listing at http://localhost:8080/")
    http.ListenAndServe(":8080", nil)
}
