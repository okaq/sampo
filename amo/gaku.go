// okaq web server
// gaku draw tool
// AQ <aq@okaq.com>
// 2018-05-06
package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "math/rand"
    "net/http"
    "strconv"
    "strings"
    "time"
)

const (
    INDEX = "gaku.html"
    // imge dir path
    IMG = "img/"
    // data dir path
    DATA = "data/"
    // backup dir path
    IMGA = "imga/"
    // json dir path
    IMGB = "imgb/"
)

var (
    // file list
    Images []string
    // pid cache
    Pids []string
    // random
    R *rand.Rand
)

func motd() {
    s0 := time.Now().String()
    fmt.Printf("OKAQ web server start\n%s\n", s0)
}

func files() {
    // read asset files into data cache
    f0, err := ioutil.ReadDir(IMG)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Printf("Found %d image files\n", len(f0))
    Images = make([]string, len(f0))
    for i,f1 := range f0 {
        // fmt.Println(f1.Name())
        Images[i] = f1.Name()
    }
    fmt.Println(Images)
}

func rng() {
    t0 := time.Now().UnixNano()
    s0 := rand.NewSource(t0)
    R = rand.New(s0)
    // test
    fmt.Sprintf("rng test max int: %d\n", R.Uint32())
}

func GakuHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w,r,INDEX)
}

func ImgHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    // static png server for import into sampler
    // fmt.Println(http.StripPrefix(r.URL))
    // fmt.Println(r.URL.String())
    // url of form "/a/pic.png"
    // fmt.Println(r.URL.String()[3:])
    s0 := fmt.Sprintf("%s%s",IMG,r.URL.String()[3:])
    // strip prefix, add dir, serve img file
    fmt.Println(s0)
    // http.ServeFile(w,r,IMG+png)
    http.ServeFile(w,r,s0)
}

func PidHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    pid := struct {
        Pid int `json:"pid"`
        Date int `json:"date"`
    } {
        0,
        0,
    }
    err := json.NewDecoder(r.Body).Decode(&pid)
    if err != nil {
        fmt.Println(err)
        // write 500 code
    }
    fmt.Println(pid)
    // not concurrency safe
    // Pids = append(Pids, pid)
    p0 := struct {
        Pid string
        Date string
        Images []string
    } {
        // rand time generators
        strconv.FormatUint(uint64(R.Uint32()),10),
        strconv.FormatInt(time.Now().UnixNano(),10),
        Images,
    }
    // s0 := fmt.Sprintf("%d:%d;%s:%s",pid.Pid,pid.Date,p0.Pid,p0.Date)
    // s1 := json.Marshal(p0)
    // w.Write([]byte(s0))
    b0, err := json.Marshal(p0)
    if err != nil {
        fmt.Println(err)
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write(b0)
}

func SaveHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    d0 := struct {
        Key string `json:"key"`
        Data []byte `json:"data"`
    } {
        "0",
        []byte{},
    }
    err := json.NewDecoder(r.Body).Decode(&d0)
    if err != nil {
        fmt.Println(err)
    }
    s0 := fmt.Sprintf("decoded %d bytes for %s image", len(d0.Data), d0.Key)
    b0 := []byte(s0)
    s1 := strings.ToLower(d0.Key)
    s2 := fmt.Sprintf("img/%s.png", s1)
    fmt.Printf("writing %s image to file %s\n", s1, s2)
    // base64 enc
    var b1 []byte
    b1, err = json.Marshal(d0.Data)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(string(b1))
    // uncompressed bits
    w.Write(b0)
}

func main() {
    motd()
    rng()
    files()
    http.HandleFunc("/", GakuHandler)
    http.HandleFunc("/a", PidHandler)
    http.HandleFunc("/b/", ImgHandler)
    http.HandleFunc("/c", SaveHandler)
    http.ListenAndServe(":8080", nil)
}


