package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/card"
	"github.com/stripe/stripe-go/charge"
	"github.com/stripe/stripe-go/customer"
)

type Plan struct {
	Amount int64  `json:"amount"`
	Id     string `json:"id"`
	Name   string `json:"name"`
}

type Sub struct {
	PeriodEnd   int64  `json:"period_end"`
	PeriodStart int64  `json:"period_start"`
	Plan        Plan   `json:"plan"`
	Quantity    int64  `json:"quantity"`
	Status      string `json:"status"`
}

func main() {
	router := httprouter.New()
	router.GET("/payment", handlePayment)
	router.GET("/source", getSource)
	router.POST("/payment", handlePayment)
	router.GET("/updateSource", updateSource)
	router.GET("/sub", getSub)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "DELETE", "PUT", "PATCH", "OPTIONS"},
	})

	log.Fatal(http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", c.Handler(router)))
}

func handlePayment(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	stripeToken := r.FormValue("stripeToken")
	stripe.Key = os.Getenv("STRIPE_KEY")

	chargeParams := &stripe.ChargeParams{
		Amount:   45000,
		Currency: "usd",
		Desc:     "Charge for test@example.com",
	}
	chargeParams.SetSource(stripeToken)
	ch, err := charge.New(chargeParams)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(ch)
}

func createCustomer() {
	// account ID
	// email, first, last name
	// card
}

func getSub(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	sub := Sub{
		Quantity:    2,
		Status:      "active",
		PeriodStart: 1472862112,
		PeriodEnd:   1475454112,
		Plan: Plan{
			Amount: 45010,
			Id:     "cass.prod.medium",
			Name:   "Cass Production Medium",
		},
	}
	subs := []Sub{sub}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(subs)
}

func createSubscription() {
	// get customer ID
	// plan based on cluster size & node size
}

func updateSource(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// get customer ID
	// get new credit cards info
	// stripeToken := r.FormValue("stripeToken")
	customerID := os.Getenv("STRIPE_CUSTOMER_ID")
	c, err := customer.Get(customerID, nil)
	if err != nil {
		panic(err)
	}
	defaultCard := c.DefaultSource
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(defaultCard)
}

func getSource(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	stripe.Key = os.Getenv("STRIPE_KEY")
	customerId := os.Getenv("STRIPE_CUSTOMER_ID")

	params := &stripe.CardListParams{Customer: customerId}
	// params.Filters.AddFilter("limit", "", "3")
	i := card.List(params)
	for i.Next() {
		c := i.Card()
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(c)
	}

}
