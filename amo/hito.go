// okaq hito web serve
// bitmap test sequence
// AQ <aq@okaq.com>
// 2018-05-22
package main

import (
    "fmt"
    "net/http"
    "time"
)

const (
    INDEX = "hito.html"
)

func HitoHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w,r,INDEX)
}

func main() {
    fmt.Printf("okaq hito start on localhost:8080\n%s\n", time.Now().String())
    // do files
    http.HandleFunc("/", HitoHandler)
    http.ListenAndServe(":8080", nil)
}


