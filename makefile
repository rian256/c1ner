BINARY_NAME=c1ner

run:
	echo "Running main.go"
	go run main.go

compile:
	echo "Compiling for every OS and Platform"
    GOOS=freebsd GOARCH=386 go build -o bin/${BINARY_NAME}-freebsd-386 main.go
    GOOS=linux GOARCH=386 go build -o bin/${BINARY_NAME}-linux-386 main.go
    GOOS=windows GOARCH=386 go build -o bin/${BINARY_NAME}-windows-386 main.go