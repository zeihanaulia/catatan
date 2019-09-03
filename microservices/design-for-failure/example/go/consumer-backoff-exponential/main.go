package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/cenkalti/backoff/v3"
	"github.com/zeihanaulia/microservice/design-for-failure/go/shared"
)

func Backoff() (string, error) {

	var body string
	start := time.Now()
	operation := func() error {
		elapsed := time.Since(start)
		log.Printf("Retry took %s", elapsed)

		var err error
		body, err = shared.CallProducerService()
		if err != nil {
			return err
		}
		return nil
	}

	bo := backoff.WithMaxRetries(backoff.NewExponentialBackOff(), 10)
	err := backoff.Retry(operation, bo)
	if err != nil {
		log.Printf("End of Retry %s", err.Error())
		return body, err
	}

	return body, nil
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		producer, err := Backoff()
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
