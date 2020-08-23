build: generate
	CGO_ENABLED=0 go build

generate: 
	go generate

run-example: generate
	go run main.go --force -f ./example/student.yaml -o ./example/student

clean:
	rm -rf ./template/**/*.qtpl.go
	rm -rf ./autoAPI

build-release: generate
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o autoAPI.exe
	zip autoAPI-windows-amd64.zip autoAPI.exe
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o autoAPI
	zip autoAPI-darwin-amd64.zip autoAPI
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o autoAPI
	zip autoAPI-linux-amd64.zip autoAPI
