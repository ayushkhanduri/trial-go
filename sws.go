package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w,"Hi there, I love %s! and other parameters are %s", r.URL.Path[1:],r.URL.Path[0:])
}

func linuxEnd(w http.ResponseWriter , r *http.Request){
  fmt.Fprintf(w,"This is the secret linux end , shhhhhhhh.....")
}
func main() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/root",linuxEnd)
    http.ListenAndServe(":8080", nil)
}
