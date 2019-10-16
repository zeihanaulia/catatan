package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/Rican7/retry"
	"github.com/Rican7/retry/backoff"
	"github.com/Rican7/retry/jitter"
	"github.com/Rican7/retry/strategy"
	"github.com/zeihanaulia/microservice/design-for-failure/go/shared"
)

func Backoff() (string, error) {

	var body string
	start := time.Now()

	action := func(attempt uint) error {
		elapsed := time.Since(start)
		log.Printf("Retry took %s", elapsed)

		var err error
		body, err = shared.CallProducerService()
		if err != nil {
			return err
		}
		return nil
	}

	seed := time.Now().UnixNano()
	random := rand.New(rand.NewSource(seed))

	err := retry.Retry(
		action,
		strategy.Limit(10),
		strategy.BackoffWithJitter(
			backoff.Exponential(500*time.Millisecond, 1.5),
			jitter.Deviation(random, 0.5),
		),
	)

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
