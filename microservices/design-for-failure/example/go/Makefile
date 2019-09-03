.PHONY: clean deploy

clean:
	rm -rf ./bin

producer: clean
	env GOOS=darwin go build -ldflags="-s -w" -o bin/producer producer/main.go
	./bin/producer 

consumer: clean
	env GOOS=darwin go build -ldflags="-s -w" -o bin/consumer consumer/main.go
	./bin/consumer 

consumer-backoff-exponential: clean
	env GOOS=darwin go build -ldflags="-s -w" -o bin/consumer-backoff-exponential consumer-backoff-exponential/main.go
	./bin/consumer-backoff-exponential

consumer-backoff-jitter: clean
	env GOOS=darwin go build -ldflags="-s -w" -o bin/consumer-backoff-jitter consumer-backoff-jitter/main.go
	./bin/consumer-backoff-jitter