// okaq hito web serve
// bitmap test sequence
// AQ <aq@okaq.com>
// 2018-05-22
package main

import (
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

func HitoHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w,r,INDEX)
}

func main() {
    fmt.Printf("okaq hito start on localhost:8080\n%s\n", time.Now().String())
    // do files
    files()
    http.HandleFunc("/", HitoHandler)
    http.ListenAndServe(":8080", nil)
}


