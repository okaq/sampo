// okaq sampo eiga
// bitmap font sampler
// 2018-04-20
package main

import (
    // "encoding/json"
    "fmt"
    "math/rand"
    "net/http"
    "strconv"
    "sync"
    "time"
)

const (
    PATH = "data/"
    JSON = "data/json"
    IMG = "data/img"
    INDEX = "eiga.html"
)

var (
    Sid string
    Rng *rand.Rand
    C *Cache
    M chan *Message
)

type Cache struct {
    S map[string]string
    *sync.Mutex
}

func NewCache() *Cache {
    c0 := &Cache{}
    c0.S = make(map[string]string)
    return c0
}

type Message struct {
    K string
    V string
}

func NewMessage() *Message {
    m0 := &Message{}
    return m0
}

func Rando() {
    t0 := time.Now().UnixNano()
    s0 := rand.NewSource(t0)
    Rng = rand.New(s0)
}

func Session() {
    // create session id and store in cache
    f0 := Rng.Uint32()
    f1 := uint64(f0)
    s0 := strconv.FormatUint(f1, 10)
    // C.S["session"] = s0
    m0 := NewMessage()
    m0.K = "session"
    m0.V = s0
    go func() { M <- m0 }()
    // channel with type key,val strings
    t0 := time.Now().UnixNano()
    s1 := strconv.FormatInt(t0, 10)
    // C.S["time"] = s1
    m1 := NewMessage()
    m1.K = "time"
    m1.V = s1
    go func() { M <- m1 }()
}

func Receiver() {
    // start the channel
    M := make(chan *Message)
    fmt.Println(M)
    go func() {
        for {
            m0 := <-M
            fmt.Println(m0)
            C.Lock()
            C.S[m0.K] = m0.V
            C.Unlock()
            fmt.Println(C)
        }
    }()
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

func PidHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    // endpoint for session id
    // 
}

func main() {
    t0 := time.Now()
    fmt.Printf("okaq eiga web serve started at %s\n", t0.String())
    Rando()
    fmt.Printf("rand test %d\n", Rng.Int())
    fmt.Printf("json path located at %s\n", JSON)
    fmt.Printf("img path located at %s\n", IMG)
    C = NewCache()
    fmt.Println(C)
    Receiver()
    Session()
    time.Sleep(10)
    fmt.Printf("session id %s\n", C.S["session"])
    fmt.Printf("session time %s\n", C.S["time"])
    // Receiver()
    http.HandleFunc("/", EigaHandler)
    http.HandleFunc("/a", PidHandler)
    http.HandleFunc("/b", ImgHandler)
    http.HandleFunc("/c", JsonHandler)
    http.ListenAndServe(":8080", nil)
}
