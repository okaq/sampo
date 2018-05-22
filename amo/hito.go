// okaq hito web serve
// bitmap test sequence
// AQ <aq@okaq.com>
// 2018-05-22
package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "time"
)

const (
    INDEX = "hito.html"
    // json file path
    IMGB = "imgb/"
)

var (
    F []string
    J []byte
)

func files() {
    F = make([]string, 0)
    f0, err := ioutil.ReadDir(IMGB)
    if err != nil {
        fmt.Println(err)
    }
    for _, f := range f0 {
        // fmt.Println(f.Name())
        // fmt.Printf("%d: %s%s\n", i, IMGB, f.Name())
        s0 := fmt.Sprintf("%s%s", IMGB, f.Name())
        F = append(F, s0)
    }
    fmt.Println(F)
}

func pack() {
    var err error
    J, err = json.Marshal(F)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(J)
}

func HitoHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w,r,INDEX)
}

func JsonHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
}

func PathHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    w.Header().Set("Content-Type", "application/json")
    w.Write(J)
}

func main() {
    fmt.Printf("okaq hito start on localhost:8080\n%s\n", time.Now().String())
    // do files
    files()
    pack()
    http.HandleFunc("/", HitoHandler)
    http.HandleFunc("/a", PathHandler)
    http.HandleFunc("/b/", JsonHandler)
    http.ListenAndServe(":8080", nil)
}


