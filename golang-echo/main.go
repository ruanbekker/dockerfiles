package main

import (
    "fmt"
    "time"
    "os"
)

func main() {

    hostname, err := os.Hostname()
    if err != nil {
	fmt.Println(err)
	os.Exit(1)
    }
    i := 0

    for {
        t := time.Now()
        i++
	fmt.Printf("%s : ping from %s : iteration %d\n", t.Format(time.RFC3339), hostname, i)
        //fmt.Println(t.Format(time.RFC3339) + ": echo :", i)
        time.Sleep(time.Second * 5)
    }

}
