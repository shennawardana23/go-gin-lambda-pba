build:
	GOOS=linux go build -o ./bootstrap ./
	zip main.zip bootstrap
run:
	go run main.go