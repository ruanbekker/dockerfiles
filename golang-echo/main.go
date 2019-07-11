package main

import (
    "fmt"
    "time"
)

func main() {

    i := 0
    for {
        t := time.Now()
        i++
        fmt.Println(t.Format(time.RFC3339) + ": echo :", i)
        time.Sleep(time.Second * 5)
    }

}
