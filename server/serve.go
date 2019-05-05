package server

import (
	"net/http"
	"log"
)

const port = ":8080"

func Start() {
	log.Println("Starting Server on port " + port + "...")

	http.Handle("/webhook", &dfReqHandler{})

	http.ListenAndServe(port, nil)
}
