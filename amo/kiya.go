// okaq web server
// "the cookie problem"
// stats, player grid
// AQ <aq@okaq.com>
// 2018-06-28
package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	INDEX = "kiya.html"
)

func KiyaHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX)
}

func motd() {
	fmt.Println("okaq cookies web serve on \"localhost:8080\"")
	fmt.Printf("%s\n", time.Now().String())
}

func main() {
	motd()
	http.HandleFunc("/", KiyaHandler)
	http.ListenAndServe(":8080", nil)
)


