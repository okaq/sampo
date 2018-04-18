// okaq sampo eiga
// bitmap font sampler
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

func Rando() {
    t0 := time.Now().UnixNano()
    s0 := rand.NewSource(t0)
    Rng = rand.New(s0)
}

func EigaHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w,r,INDEX)
}

func main() {
    t0 := time.Now()
    fmt.Printf("okaq eiga web serve started at %s\n", t0.String())
    Rando()
    fmt.Printf("session id %d\n", Rng.Int())
    http.HandleFunc("/", EigaHandler)
    http.ListenAndServe(":8080", nil)
}
