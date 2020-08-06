build: generate
	go build

generate: 
	go generate

run-example: generate
	go run main.go -f ./example/student.yaml -o ./example/student

clean:
	rm -rf ./template/**/*.qtpl.go
	rm -rf ./autoAPI

build-release: generate
	GOOS=windows GOARCH=amd64 go build -o autoAPI-windows-amd64.exe
	GOOS=darwin GOARCH=amd64 go build -o autoAPI-darwin-amd64
	GOOS=linux GOARCH=amd64 go build -o autoAPI-linux-amd64
