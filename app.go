package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func main() {
    http.HandleFunc("/", hello)
    err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
    if err != nil {
        log.Fatal("ListenAndServe:", err)
    }
}

func hello(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintln(w, "<h1>hello, world!</h1>")
    fmt.Fprintln(w, "<p>This instance number is "+os.Getenv("CF_INSTANCE_INDEX")+"</p>")
}
