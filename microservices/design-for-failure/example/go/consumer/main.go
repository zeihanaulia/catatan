package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/zeihanaulia/microservice/design-for-failure/go/shared"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		producer, err := shared.CallProducerService()
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprintf(w, fmt.Sprintf("{\"error\":\"%v\"}", err))
			return
		}
		fmt.Fprintf(w, fmt.Sprintf("{\"message\":\"%v\"}", producer))
	})
	log.Println("listening at :8091")
	_ = http.ListenAndServe(":8091", nil)
}
