// okaq web server
// "cookies"
// AQ <aq@okaq.com>
// 2018-07-18
package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	INDEX = "lima.html"
)

func LimaHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServerFile(w,r,INDEX)
}

func motd() {
	fmt.Println("okaq cookies started localhost:8080")
	fmt.Printf("start time: %s\n", time.Now().String())
}

func main() {
	motd()
	http.HandleFunc("/", LimaHandler)
	http.ListenAndServe(":8080", nil)
}


