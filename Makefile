build: generate
	go build

generate: 
	go generate

run-example: generate
	go run main.go -f ./example/student.yaml -o ./example/student