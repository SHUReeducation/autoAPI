build: generate
	go build
	chmod +x ./autoAPI

generate: 
	go generate
