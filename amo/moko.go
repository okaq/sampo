// okaq web
// gltf load
// AQ <aq@okaq.com>
// 2018-08-06
package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	INDEX = "moko.html"
)

func MokoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX)
}

func motd() {
	fmt.Println("okaq gltf started localhost:8080")
	fmt.Printf("start: %s\n", time.Now().String())
}

func main() {
	motd()
	http.HandleFunc("/", MokoHandler)
	http.ListenAndServe(":8080", nil)
}


