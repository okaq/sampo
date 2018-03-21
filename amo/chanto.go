// okaq web serve
// AQ <aq@okaq.com>
// 2018-03-20
package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "sync"
    "sync/atomic"
)

const (
    INDEX = "chanto.html"
)

var (
    Counter uint64
    Cach *Cache
)

type Cache struct {
    C map[string]string
    *sync.Mutex
}

func NewCache() *Cache {
    c0 := &Cache{}
    c0.C = make(map[string]string)
    return c0
}

func ChanHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    // increment counter
    _ = atomic.AddUint64(&Counter, 1)
    http.ServeFile(w, r, INDEX)
}

func StatHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    // pipeline atomic counter
    c0 := atomic.LoadUint64(&Counter)
    // json response
    c1 := struct {
        Count uint64 `json:"count"`
    } {
        c0,
    }
    j0, err := json.Marshal(c1)
    if err != nil {
        fmt.Println(err)
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write(j0)
}

func Count() {
    // init counter
    Counter = 0
}

func main() {
    fmt.Println("web serving on localhost:8080")
    Count()
    Cach = NewCache()
    http.HandleFunc("/", ChanHandler)
    http.HandleFunc("/s", StatHandler)
    http.ListenAndServe(":8080", nil)
}


