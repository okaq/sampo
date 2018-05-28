// okaq web server
// "the cookie problem"
// AQ <aq@okaq.com>
// 2018-05-30
package main

import (
    "fmt"
    "net/http"
    "time"
)

const (
    INDEX = "jiyu.html"
)

func motd() {
    fmt.Println("okaq starting on localhost:8080")
    fmt.Printf("%s\n", time.Now().ToString())
}

func JiyuHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w,r,INDEX)
}

func StatHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    // accumulator
}

func main() {
    motd()
    http.HandleFunc("/", JiyuHandler)
    http.HandleFunc("/a", StatHandler)
    http.ListenAndServe(":8080", nil)
}


