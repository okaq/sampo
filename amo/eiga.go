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
    JSON = "data/json"
    IMG = "data/img"
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

func ImgHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    // endpoint for bitmap storage
}

func JsonHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    // endpoint for json storage
}

func main() {
    t0 := time.Now()
    fmt.Printf("okaq eiga web serve started at %s\n", t0.String())
    Rando()
    fmt.Printf("session id %d\n", Rng.Int())
    fmt.Printf("json path located at %s\n", JSON)
    fmt.Printf("img path located at %s\n", IMG)
    http.HandleFunc("/", EigaHandler)
    http.HandleFunc("/a", ImgHandler)
    http.HandleFunc("/b", JsonHandler)
    http.ListenAndServe(":8080", nil)
}
