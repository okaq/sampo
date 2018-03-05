// okaq amo
// AQ <aq@okaq.com>
// 2018-03-02
package main

import (
    "fmt"
    "net/http"
    "os"
)

const (
    // try open file for writing
    // create new if it doesnt exist
    CFG = "cfg.txt"
    SRT = "srt.txt"
    PAT = "data/"
)

var (
    Scratch *os.File
    Config *os.File
)

func Load() {
    // global state
    p0 := fmt.Sprintf("%s%s", PAT, CFG)
    p1 := fmt.Sprintf("%s%s", PAT, SRT)
    fmt.Println(p0, p1)
}

func Handler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
}

func main() {
    fmt.Println("starting amo web")
    fmt.Println("opening data files")
    Load()
    // server
    // stats handler
    // config load
    // scratch load
}


