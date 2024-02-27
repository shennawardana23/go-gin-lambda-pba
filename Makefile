env=prod
mport=8080

build:
	GOOS=linux go build -o ./bootstrap ./
	zip main.zip bootstrap

run:
	go run main.go ENV=${env} MPORT=${mport}
