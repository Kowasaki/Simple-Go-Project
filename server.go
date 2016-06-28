package main

import (
    "net/http"
)

func main() {
    http.HandleFunc("/", testmsg)
    http.ListenAndServe(":5006", nil)
}

func testmsg(w http.ResponseWriter, r *http.Request)  {
    w.Write([]byte("Hey There!"))

}
