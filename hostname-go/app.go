package main

import (
    "fmt"
    "os"
    "log"
    "net/http"
)

func hostnameHandler(w http.ResponseWriter, r *http.Request) {
    myhostname, _ := os.Hostname()
    log.Println("logging request for hostnameHandler on /")
    fmt.Fprintln(w, "Hostname:", myhostname)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "ok")
}

func main() {
    const port string = "8080"
    log.Println("Server listening on port", port)
    http.HandleFunc("/", hostnameHandler)
    http.HandleFunc("/health", healthHandler)
    http.ListenAndServe(":" + port, nil)
}
