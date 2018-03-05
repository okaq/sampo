// okaq amo
// AQ <aq@okaq.com>
// 2018-03-02
package main

import (
    "fmt"
    "net/http"
)

const (
    // try open file for writing
    // create new if it doesnt exist
    CFG = "cfg.txt"
    SRT = "sat.txt"
    PAT = "data/"
)

func Handler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
}

func main() {
    fmt.Println("starting amo web")
    fmt.Println("opening data files")
    // server
    // stats handler
    // config load
    // scratch load
}
