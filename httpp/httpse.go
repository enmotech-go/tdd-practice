package main

import (
	"fmt"
	"log"
	"net/http"
)

func PlayerServer(recorder http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(recorder, "20")
}

func main() {
	handler := http.HandlerFunc(PlayerServer)
	if err := http.ListenAndServe(":5000", handler); err != nil {
		log.Fatalf("can not server on 5000 %v" ,err)
	}
}
