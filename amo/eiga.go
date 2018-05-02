// okaq sampo eiga
// bitmap font sampler
// 2018-04-20
package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "math/rand"
    "net/http"
    "strconv"
    "sync"
    "sync/atomic"
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
    // atomic pid counter
    N uint64
    // chronological keys
    K []string
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
    // go func() { M <- m0 }()
    // channel with type key,val strings
    t0 := time.Now().UnixNano()
    s1 := strconv.FormatInt(t0, 10)
    // C.S["time"] = s1
    m1 := NewMessage()
    m1.K = "time"
    m1.V = s1
    go func() {
        M <- m0
        M <- m1
    }()
}

func Receiver() {
    // start the channel
    M = make(chan *Message)
    // fmt.Println(M)
    go func() {
        for {
            m0 := <-M
            // fmt.Println(m0.K)
            // C.Lock()
            C.S[m0.K] = m0.V
            // C.Unlock()
            // fmt.Println(C)
            // mutex lock or chan
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
    d0 := json.NewDecoder(r.Body)
    var j0 map[string][]byte
    var err error
    var b0 []byte
    err = d0.Decode(&j0)
    if err != nil {
        fmt.Println(err)
    }
    b0, err = json.Marshal(j0)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(string(b0))
    // fmt.Println(j0["0"])
    // ok, byte array encoded as base64 string
    // stich via format to json for disk save
    f0 := fmt.Sprintf("%s/%s", JSON, "gugi.json")
    err = ioutil.WriteFile(f0, b0, 0644)
    if err != nil {
        fmt.Println(err)
    }
    // w.Write([]byte("Json bitmap data decoded"))
    w.Write(b0)
}

func PidHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    // endpoint for session id
    p0 := struct {
        Time uint64 `json:"time"`
        Session uint64 `json:"session"`
    } {
        0,
        0,
    }
    err := json.NewDecoder(r.Body).Decode(&p0)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(p0)
    s0 := strconv.FormatUint(p0.Time, 10)
    s1 := strconv.FormatUint(p0.Session, 10)
    // m0 := NewMessage()
    atomic.AddUint64(&N, 1)
    // update pids keys list
    // server session id
    f0 := Rng.Uint32()
    f1 := uint64(f0)
    s2 := strconv.FormatUint(f1, 10)
    // server time stamp
    t0 := time.Now().UnixNano()
    s3 := strconv.FormatInt(t0, 10)
    s4 := fmt.Sprintf("%s:%s;%s:%s", s2, s3, s1, s0)
    s5 := atomic.LoadUint64(&N)
    fmt.Printf("pid key for user # %d: %s\n", s5, s4)
    // set the key, synchronously
    K = append(K, s4)
    fmt.Println(K)
    // send as text, todo: marshal json
    w.Header().Set("Content-Type", "text/plain")
    w.Write([]byte(s4))

    // coordinate pid schema
    // client id, server id
    // each pid of form [time,session]
    // abstract to key, value cache on server
    // maintain singleton state for pid #1
}

func main() {
    t0 := time.Now()
    fmt.Printf("okaq eiga web serve started at %s\n", t0.String())
    N = 0
    fmt.Printf("atomic counter init with server process = %d\n", N)
    Rando()
    s0 := Rng.Int()
    fmt.Printf("rand test %d\n", s0)
    fmt.Printf("json path located at %s\n", JSON)
    fmt.Printf("img path located at %s\n", IMG)
    C = NewCache()
    // fmt.Println(C)
    Receiver()
    Session()
    time.Sleep(100*time.Millisecond)
    fmt.Printf("session id %s\n", C.S["session"])
    fmt.Printf("session time %s\n", C.S["time"])
    K = make([]string,0)
    s1 := fmt.Sprintf("%d:%d",s0,t0.UnixNano())
    K = append(K, s1)
    fmt.Println(K)
    // Receiver()
    http.HandleFunc("/", EigaHandler)
    http.HandleFunc("/a", PidHandler)
    http.HandleFunc("/b", ImgHandler)
    http.HandleFunc("/c", JsonHandler)
    http.ListenAndServe(":8080", nil)
}


