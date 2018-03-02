// okaq amo
// AQ <aq@okaq.com>
// 2018-03-02
package main

import (
    "fmt"
    "net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
}

func main() {
    fmt.Println("starting amo web")
    // server
}
