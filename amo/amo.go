// okaq amo
// AQ <aq@okaq.com>
// 2018-03-02
package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "os"
    "sync/atomic"
)

const (
    // try open file for writing
    // create new if it doesnt exist
    CFG = "cfg.txt"
    SRT = "srt.txt"
    PAT = "data/"
    // web file
    INDEX = "amo.html"
)

var (
    Scratch *os.File
    Config *os.File
    Counter uint64
)

func Load() {
    var err error
    // global state
    p0 := fmt.Sprintf("%s%s", PAT, CFG)
    p1 := fmt.Sprintf("%s%s", PAT, SRT)
    fmt.Println(p0, p1)
    // config open
    Config, err = os.Open(p0)
    if err != nil {
        fmt.Println(err)
    }
    // scratch open
    Scratch, err = os.Open(p1)
    if err != nil {
        fmt.Println(err)
    }
}

func AmoHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    atomic.AddUint64(&Counter, 1)
    http.ServeFile(w, r, INDEX)
}

func StatHandler(w http.ResponseWriter, r *http.Request) {
    // ouput statistics
    // to web client
    fmt.Println(r)
    // golang anonymous struct
    stat := struct {
        Count uint64 `json:"count"`
        ID string `json:"id"`
    } {
        atomic.LoadUint64(&Counter),
        "amo_server",
    }
    js1, err := json.Marshal(stat)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(stat)
    fmt.Println(js1)
}

func main() {
    fmt.Println("starting amo web")
    fmt.Println("opening data files")
    Load()
    // server
    // stats handler
    // config load
    // scratch load
    http.HandleFunc("/", AmoHandler)
    http.HandleFunc("/s", StatHandler)
    http.ListenAndServe(":8080", nil)
}


