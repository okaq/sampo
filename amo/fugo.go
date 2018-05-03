// okaq web server
// gugi bitmap font render test
// AQ <aq@okaq.com>
// 2018-05-04
package main

import (
    "fmt"
    "net/http"
    "time"
)

var (
    INDEX = "fugo.html"
)

func motd() {
    s0 := time.Now().String()
    fmt.Printf("OKAQ Web Server started\n%s\n", s0)
}

func FugoHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w,r,INDEX)
}

func main() {
    motd()
    http.HandleFunc("/", FugoHandler)
    http.ListenAndServe(":8080", nil)
}


