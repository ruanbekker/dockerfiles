package main

import (
    "fmt"
    "os"
    "net/http"
    "encoding/json"
)

type User struct  {
    Name    string `json:"name"`
    Twitter string `json:"twitter"`
    Contact string `json:"contact"`
    Blogs []string `json:"blogs"`
}

func main() {
    const port string = "8000"
    fmt.Println("Server listening on port", port)
    http.HandleFunc("/", hostnameHandler)
    http.HandleFunc("/about", aboutHandler)
    http.ListenAndServe(":" + port, nil)
}

func hostnameHandler(w http.ResponseWriter, r *http.Request) {
    myhostname, _ := os.Hostname()
    fmt.Fprintln(w, "Hostname:", myhostname)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    user := User{
        Name: "Ruan",
        Twitter: "@ruanbekker",
        Contact: "https://ruan.dev",
        Blogs: []string{"https://blog.ruanbekker.com", "https://sysadmins.co.za", "https://ruan.dev/blog"},
    }
    if err := json.NewEncoder(w).Encode(user); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}
