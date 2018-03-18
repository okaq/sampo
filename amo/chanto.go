// okaq web serve
// AQ <aq@okaq.com>
// 2018-03-20
package main

import (
    "fmt"
    "net/http"
)

const (
    INDEX = "chanto.html"
)

func ChanHander(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w, r, INDEX)
}

func StatHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    // pipeline atomic counter
}

func main() {
    fmt.Println("web serving on localhost:8080")
    http.HandleFunc("/", ChanHandler)
    http.HandleFunc("/s", StatHandler)
    http.ListenAndServe(":8080", nil)
}


