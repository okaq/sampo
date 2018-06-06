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
    DRAGON = "dragon.html"
)

func motd() {
    fmt.Println("okaq starting on localhost:8080")
    fmt.Printf("%s\n", time.Now().String())
}

func JiyuHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w,r,INDEX)
}

func DragonHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w,r,DRAGON)
}

func StatHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    // accumulator
}

func JarHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    // jars kept in global state on server
    // session key on client can be used to decrypt
}

func main() {
    motd()
    http.HandleFunc("/", JiyuHandler)
    http.HandleFunc("/a", StatHandler)
    http.HandleFunc("/b", JarHandler)
    http.HandleFunc("/c", DragonHandler)
    http.ListenAndServe(":8080", nil)
}


