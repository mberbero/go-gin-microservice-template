swag:
	swag init -g app/cmd/main.go  

run:
	go build -o servicename app/cmd/main.go && ./servicename