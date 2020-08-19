build: generate
	go build

generate: 
	go generate

run-example: generate
	go run main.go --force -f ./example/student.yaml -o ./example/student

clean:
	rm -rf ./template/**/*.qtpl.go
	rm -rf ./autoAPI

build-release: generate
	GOOS=windows GOARCH=amd64 go build -o autoAPI.exe
	zip autoAPI-windows-amd64.zip autoAPI.exe
	GOOS=darwin GOARCH=amd64 go build -o autoAPI
	zip autoAPI-darwin-amd64.zip autoAPI
	GOOS=linux GOARCH=amd64 go build -o autoAPI
	zip autoAPI-linux-amd64.zip autoAPI
