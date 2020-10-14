package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/paymentintent"
)

func (s *Server) handleIntent() http.HandlerFunc {
	type intentRequest struct {
		Amount   int64 `json:"amount"`
		Invoices []struct {
			Number string `json:"number"`
			Amount string `json:"amount"`
		} `json:"invoices"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			s.handleError(w, err.Error())
			return
		}
		s.log(string(body))
		var bodyJSON intentRequest
		err = json.Unmarshal(body, &bodyJSON)
		if err != nil {
			s.handleError(w, err.Error())
			return
		}
		s.log(bodyJSON)

		stripe.Key = s.configuration.stripePrivateKey

		params := &stripe.PaymentIntentParams{
			Amount:   stripe.Int64(bodyJSON.Amount),
			Currency: stripe.String(string(stripe.CurrencyUSD)),
		}

		for _, invoice := range bodyJSON.Invoices {
			params.AddMetadata("Invoice "+invoice.Number, invoice.Amount)
		}
		pi, err := paymentintent.New(params)
		if err != nil {
			s.handleError(w, err.Error())
			return
		}

		output := map[string]interface{}{
			"intentID": pi.ClientSecret,
		}
		bytes, err := json.Marshal(output)
		if err != nil {
			s.handleError(w, err.Error())
			return
		}
		w.Header().Add("Content-Type", "application-json")
		w.Write(bytes)
	}
}
