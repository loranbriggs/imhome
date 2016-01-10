package main

import (
    "fmt"
    "github.com/loranbriggs/w1thermsensor"
    "html/template"
    "net/http"
    "strings"
)

var users = map[string]string{}

func rootHandler(w http.ResponseWriter, r *http.Request) {
    ip := lookupIp(r)
    _, returning := users[ip] // check map for visitor
    if returning {
        http.Redirect(w, r, "/home", http.StatusFound)
    }
    http.Redirect(w, r, "/new", http.StatusFound)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    ip := lookupIp(r)
    name := users[ip]

    data := struct {
        User string
        Users map[string]string
        Temperature string
    } {
        name,
        users,
        w1thermsensor.Temperatrue(),
    }

    t, _ := template.ParseFiles("views/home.html")
    t.Execute(w, data)
}

func newHandler(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("views/new.html")
    t.Execute(w, users)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
    ip := lookupIp(r)
    name := r.FormValue("name")
    users[ip] = name
    http.Redirect(w, r, "/", http.StatusFound)
}

func lookupIp(r *http.Request) string {
    return strings.Split(r.RemoteAddr, ":")[0]
}

func main() {
    http.HandleFunc("/", rootHandler)
    http.HandleFunc("/home", homeHandler)
    http.HandleFunc("/new", newHandler)
    http.HandleFunc("/save", saveHandler)
    fmt.Println("listing at http://localhost:8080/")
    http.ListenAndServe(":8080", nil)
}
