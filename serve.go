package main

import (
        "net/http"
        "log"
)

func main() {
        http.Handle("/", http.FileServer(http.Dir("./files")))
        log.Fatal(http.ListenAndServe(":80", nil))
}
