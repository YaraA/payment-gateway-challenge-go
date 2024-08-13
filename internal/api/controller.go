package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cko-recruitment/payment-gateway-challenge-go/docs"
	"github.com/cko-recruitment/payment-gateway-challenge-go/internal/handlers"
	paymentProcessor "github.com/cko-recruitment/payment-gateway-challenge-go/third_party/payment_processor"
	httpSwagger "github.com/swaggo/http-swagger"
)

type pong struct {
	Message string `json:"message"`
}

// PingHandler returns an http.HandlerFunc that handles HTTP Ping GET requests.
func (a *Api) PingHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(pong{Message: "pong"}); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

// SwaggerHandler returns an http.HandlerFunc that handles HTTP Swagger related requests.
func (a *Api) SwaggerHandler() http.HandlerFunc {
	return httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("http://%s/swagger/doc.json", docs.SwaggerInfo.Host)),
	)
}

// GetPaymentHandler returns an http.HandlerFunc that handles Payments GET requests.
func (a *Api) GetPaymentHandler() http.HandlerFunc {
	h := handlers.NewPaymentsHandler(a.paymentsRepo, nil)

	return h.GetHandler()
}

// PostPaymentHandler returns an http.HandlerFunc that handles Payments Post requests.
func (a *Api) PostPaymentHandler() http.HandlerFunc {
	h := handlers.NewPaymentsHandler(a.paymentsRepo, paymentProcessor.GetBankPaymentProcessor())
	return h.PostHandler()
}
