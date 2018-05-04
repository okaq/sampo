// okaq web server
// gaku draw tool
// AQ <aq@okaq.com>
// 2018-05-06
package main

import (
    "fmt"
    "net/http"
    "time"
)

const (
    INDEX = "gaku.html"
)

func motd() {
    s0 := time.Now().String()
    fmt.Printf("OKAQ web server start\n%s\n", s0)
}

func GakuHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w,r,INDEX)
}

func main() {
    motd()
    http.HandleFunc("/", GakuHandler)
    http.ListenAndServe(":8080", nil)
}


