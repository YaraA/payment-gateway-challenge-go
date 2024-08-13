This is a Payment Gateway service. That processes payment requests.
## How to run the payment gateway
We have two options to run our project. 
### 1. Using Docker Compose
- Run the command `docker-compose up --build`
### 2. Locally on your machine
- Set the env vars `export ENV="dev" PORT="8090" BANK_PAYMENT_PROCESSOR_BASE_URL="http://localhost:8080"` 
- Run `go run main.go`

## Template structure
```
main.go - a skeleton Payment Gateway API
imposters/ - contains the bank simulator configuration.
docs/docs.go - Generated file by Swaggo
internal/handlers - contains the request handlers
internal/validator - contains validotor implementation and custom validation rules
third_party/payment_processor - contains the external service (paymentProcessor) requests 
docker-compose.yml - configures the bank simulator and the payment service
.editorconfig - It ensures a consistent set of rules for submissions when reformatting code
.goreleaser.yml - Goreleaser configuration
```

## API Documentation
http://localhost:8090/swagger/index.html.
