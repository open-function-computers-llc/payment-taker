package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"gopkg.in/gomail.v2"
)

func (s *Server) handleReceipt() http.HandlerFunc {
	type incomingPayload struct {
		Email string `json:"email"`
		Total string `json:"total"`
		Fee   string `json:"fee"`
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

		go func() {
			s.log("Sending customer email...")
			m := gomail.NewMessage()
			m.SetHeader("From", s.configuration.emailFrom)
			m.SetHeader("To", incomingData.Email)
			m.SetHeader("Subject", "Payment Reciept")
			m.SetBody("text/html", strings.ReplaceAll(receiptHTML, "%%TOTAL%%", incomingData.Total))

			port := s.configuration.smtpPort
			d := gomail.NewDialer(os.Getenv("SMTP_HOST"),
				port,
				s.configuration.smtpUser,
				s.configuration.smtpPass)
			if err := d.DialAndSend(m); err != nil {
				s.logger.Error(err)
			}
		}()
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
		// }()

		w.Write([]byte("done!"))
	}
}

var receiptHTML string = `
<h1>Thank you for your payment of $%%TOTAL%%!</h1>
`
