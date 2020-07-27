package main

import (
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	// Declare a new router
	r := mux.NewRouter()
	var err error
	r.HandleFunc("/hello", handler).Methods("GET")
	err = http.ListenAndServe(":9000", r)

	if err != nil {
		log.Errorln("Server Listening Error",err)
		return
	}

}

func handler(w http.ResponseWriter, _ *http.Request) {
	_, err := fmt.Fprintf(w, "Hello, World")
	if err != nil {
		log.Errorln("Sending Response Error",err)
	}
}
