package main

import (
	"fmt"
	"net/http"
	"time"
)

func currentTimeHandler(w http.ResponseWriter, req *http.Request) {
	currentTime := time.Now().Local()

	fmt.Fprint(w, currentTime)
}

func main() {
	http.HandleFunc("/api/timestamp", currentTimeHandler)

	http.ListenAndServe(":9999", nil)
}
