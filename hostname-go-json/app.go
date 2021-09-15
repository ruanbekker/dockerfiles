package main

import (
	"fmt"
	"os"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.Info("starting server")
	http.HandleFunc("/", hostnameHandler)
	http.ListenAndServe("0.0.0.0:8080", nil)
}

func hostnameHandler(w http.ResponseWriter, r *http.Request) {
	myhostname, _ := os.Hostname()
	log.SetFormatter(&log.JSONFormatter{})
	log.WithFields(log.Fields{"hostname": myhostname}).Info("fetched your hostname for you")
	fmt.Fprintln(w, "Hostname:", myhostname)
}
