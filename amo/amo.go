// okaq amo
// AQ <aq@okaq.com>
// 2018-03-02
package main

import (
    "encoding/json"
    "fmt"
    "math/rand"
    "net/http"
    "os"
    "strconv"
    "sync"
    "sync/atomic"
    "time"
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
    Random *rand.Rand
    Cache1 map[string]string
    Cach *Cache
    Chan chan string
)

type Cache struct {
    C map[string]string
    *sync.Mutex
}

func NewCache() *Cache {
    return &Cache{make(map[string]string),&sync.Mutex{}}
}

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

func Rng() {
    // setup the rand
    s0 := rand.NewSource(time.Now().UnixNano())
    Random = rand.New(s0)
}

func Store() {
    // make the cache
    Cache1 = make(map[string]string)
    Cache1["zero"] = "hero"
    fmt.Println(Cache1)
    // needs channel to sync access
}

func Store2() {
    // create the cache
    // launch goroutine for reciever
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
    w.Header().Set("Content-Type","application/json")
    w.Write(js1)
}

func PidHandler(w http.ResponseWriter, r *http.Request) {
    // player cache
    // pid formed from browser id, timestamp
    // and server id, timestamp combo
    // decoder
    // browser side player id
    pid := struct {
        Id string `json:"id"`
        Time string `json:"time"`
    } {
        "0",
        "0",
    }
    err := json.NewDecoder(r.Body).Decode(&pid)
    if err != nil {
        // write 500
        fmt.Println(err)
    }
    fmt.Println(pid)
    // server side player id
    sid := struct {
        Id uint64 `json:"id"`
        Time int64 `json:"time"`
    } {
        Random.Uint64(),
        time.Now().UnixNano(),
    }
    fmt.Println(sid)
    // put it all together
    tid := struct {
        Pid string `json:"pid"`
        Ptime string `json:"ptime"`
        Sid string `json:"Sid"`
        Stime string `json:"stime"`
    } {
        pid.Id,
        pid.Time,
        strconv.FormatInt(int64(sid.Id), 10),
        strconv.FormatInt(sid.Time, 10),
    }
    fmt.Println(tid)
    // store in server side cache
    // log ip, user agent for analysis
    js0, err := json.Marshal(tid)
    if err != nil {
        fmt.Println(err)
        // log and write 500
    }
    fmt.Println(js0)
    w.Header().Set("Content-Type", "application/json")
    w.Write(js0)
}

func GameHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    // game state update and response
    // struct for current state of play
}

func main() {
    fmt.Println("starting amo web")
    fmt.Println("opening data files")
    Load()
    Rng()
    Store()
    Cach = NewCache()
    fmt.Println(Cach)
    // server
    // stats handler
    // config load
    // scratch load
    http.HandleFunc("/", AmoHandler)
    http.HandleFunc("/s", StatHandler)
    http.HandleFunc("/p", PidHandler)
    http.HandleFunc("/q", GameHandler)
    http.ListenAndServe(":8080", nil)
}


