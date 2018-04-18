// okaq sampo eiga
// bitmap font smapler
// 2018-04-20
package main

import (
    "fmt"
    "math/rand"
    "net/http"
    "time"
)

const (
    PATH = "data/"
    INDEX = "eiga.html"
)

var (
    Rng *rand.Rand
)

func main() {
    t0 := time.Now()
    fmt.Println("okaq eiga web serve started at %s", t0.ToString())
    Rando()
    http.HandleFunc("/", EigaHandler)
    http.ListenAndServe(":8080", nil)
}
