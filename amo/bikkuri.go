// okaq web serve
// AQ <aq@okaq.com>
// 2018-03-14
package main

import (
    "fmt"
    "net/http"
)

const (
    INDEX = "bikkuri.html"
)

func BikHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w, r, INDEX)
}

func main() {
    fmt.Println("starting web server on localhost:8080")
    http.HandleFunc("/", BikHandler)
    http.ListenAndServe(":8080", nil)
}
