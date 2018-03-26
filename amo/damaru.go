// okaq web serve
// 2018-03-28
// AQ <aq@okaq.com>
package main

import (
    "fmt"
    "net/http"
)

const (
    INDEX = "damaru.html"
)

func DamaruHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w, r, INDEX)
}

func main() {
    fmt.Println("web serve on localhost:8080")
    http.HandleFunc("/", DamaruHandler)
    http.ListenAndServe(":8080", nil)
}


