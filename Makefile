local:
	go run cmd/local/main.go

go-zip:
	bash generate-function.sh

go-sqs:
	bash generate-function_sqs.sh