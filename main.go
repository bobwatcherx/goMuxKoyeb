package main

import (
    "fmt"
    "net/http"
    "os"
    "log"
    "github.com/gorilla/mux"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
    port = "8080"
    }

    r := mux.NewRouter()
    r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, World!")
    })
    println(port)
    err := http.ListenAndServe(":"+port, r)
    if err != nil {
    log.Fatal("ListenAndServe: ", err)
    }
}
