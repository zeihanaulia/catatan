package shared

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Response struct {
	Message string `json:"message,omitempty"`
}

func CallProducerService() (string, error) {
	read, err := http.Get("http://localhost:8090")
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(read.Body)
	if err != nil {
		return "", err
	}

	var s = new(Response)
	err = json.Unmarshal(body, &s)
	if err != nil {
		return "", err
	}

	return s.Message, nil
}
