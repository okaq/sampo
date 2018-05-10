// okaq web server
// gaku draw tool
// AQ <aq@okaq.com>
// 2018-05-06
package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "time"
)

const (
    INDEX = "gaku.html"
    // imge dir path
    IMG = "img/"
    // data dir path
    DATA = "data/"
)

var (
    // file list
    Images []string
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

func main() {
    motd()
    files()
    http.HandleFunc("/", GakuHandler)
    http.HandleFunc("/a/", ImgHandler)
    http.ListenAndServe(":8080", nil)
}


