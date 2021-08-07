package server

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func (s *Server) handleReceipt() http.HandlerFunc {
	type invoice struct {
		Number string `json:"number"`
		Amount string `json:"amount"`
	}

	type incomingPayload struct {
		Email    string `json:"email"`
		Amount   string `json:"amount"`
		Total    string `json:"total"`
		Fee      string `json:"fee"`
		Company  string `json:"company"`
		Invoices []invoice
	}

	return func(w http.ResponseWriter, r *http.Request) {
		incomingData := incomingPayload{}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			s.handleError(w, "Reading body"+err.Error())
			return
		}
		err = json.Unmarshal(body, &incomingData)
		if err != nil {
			s.handleError(w, "Invalid JSON: "+err.Error())
			return
		}
		// send a webhook notification to mattermost
		go func() {
			endpointURI := os.Getenv("MATTERMOST_NOTIFICATION_URL")
			if endpointURI == "" {
				return
			}

			s.log("Sending mattermost notification")
			messageText := "Payment completed.\nCompany: %%COMPANY%%\nEmail: %%EMAIL%%\nInvoices: %%INVOICES%%\nTotal: $%%TOTAL%%"
			messageText = strings.ReplaceAll(messageText, "%%TOTAL%%", incomingData.Amount)
			messageText = strings.ReplaceAll(messageText, "%%COMPANY%%", incomingData.Company)
			messageText = strings.ReplaceAll(messageText, "%%EMAIL%%", incomingData.Email)
			invoices := []string{}
			for _, i := range incomingData.Invoices {
				invoices = append(invoices, i.Number)
			}
			messageText = strings.ReplaceAll(messageText, "%%INVOICES%%", strings.Join(invoices, ", "))

			postBody := map[string]string{
				"text": messageText,
			}
			bodyJSON, _ := json.Marshal(postBody)
			http.Post(endpointURI, "application/json", bytes.NewBuffer(bodyJSON))
		}()

		// send an email reciept to the customer
		go func() {
			s.log("Sending customer email...")
			return // remove this line to actually send the email notification
			// m := gomail.NewMessage()
			// m.SetHeader("From", os.Getenv("EMAIL_FROM"))
			// m.SetHeader("To", os.Getenv("EMAIL_TO"))
			// m.SetHeader("Subject", s.URL+" is back up")
			// m.SetBody("text/html", "<h1>"+s.URL+" is back online!</h1><p>It was down for "+strconv.Itoa(secondsDown)+" seconds.</p>")

			// 	port := os.Getenv("SMTP_PORT")
			// 	portInt, _ := strconv.Atoi(port)
			// 	d := gomail.NewDialer(os.Getenv("SMTP_HOST"),
			// 		portInt,
			// 		os.Getenv("SMTP_USER"),
			// 		os.Getenv("SMTP_PASSWORD"))
			// 	if err := d.DialAndSend(m); err != nil {
			// 		s.Logger.Error(err)
			// 	}
			// }()
			// go func() {
			// 	s.log("Sending ofco email...")
			// 	m := gomail.NewMessage()
			// 	m.SetHeader("From", os.Getenv("EMAIL_FROM"))
			// 	m.SetHeader("To", os.Getenv("EMAIL_TO"))
			// 	m.SetHeader("Subject", s.URL+" is back up")
			// 	m.SetBody("text/html", "<h1>"+s.URL+" is back online!</h1><p>It was down for "+strconv.Itoa(secondsDown)+" seconds.</p>")

			// 	port := os.Getenv("SMTP_PORT")
			// 	portInt, _ := strconv.Atoi(port)
			// 	d := gomail.NewDialer(os.Getenv("SMTP_HOST"),
			// 		portInt,
			// 		os.Getenv("SMTP_USER"),
			// 		os.Getenv("SMTP_PASSWORD"))
			// 	if err := d.DialAndSend(m); err != nil {
			// 		s.Logger.Error(err)
			// 	}
		}()

		w.Write([]byte("done!"))
	}
}
