package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(503)
		fmt.Fprintf(w, `{"message":"Okay so your other producers function also executed successfully!"}`)
	})
	log.Println("listening at :8090")
	_ = http.ListenAndServe(":8090", nil)
}
