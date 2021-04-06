package server

import (
	"net/http"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

func (s *Server) handleReceipt() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		go func() {
			s.log("Sending customer email...")
			m := gomail.NewMessage()
			m.SetHeader("From", os.Getenv("EMAIL_FROM"))
			m.SetHeader("To", os.Getenv("EMAIL_TO"))
			m.SetHeader("Subject", s.URL+" is back up")
			m.SetBody("text/html", "<h1>"+s.URL+" is back online!</h1><p>It was down for "+strconv.Itoa(secondsDown)+" seconds.</p>")

			port := os.Getenv("SMTP_PORT")
			portInt, _ := strconv.Atoi(port)
			d := gomail.NewDialer(os.Getenv("SMTP_HOST"),
				portInt,
				os.Getenv("SMTP_USER"),
				os.Getenv("SMTP_PASSWORD"))
			if err := d.DialAndSend(m); err != nil {
				s.Logger.Error(err)
			}
		}()
		go func() {
			s.log("Sending ofco email...")
			m := gomail.NewMessage()
			m.SetHeader("From", os.Getenv("EMAIL_FROM"))
			m.SetHeader("To", os.Getenv("EMAIL_TO"))
			m.SetHeader("Subject", s.URL+" is back up")
			m.SetBody("text/html", "<h1>"+s.URL+" is back online!</h1><p>It was down for "+strconv.Itoa(secondsDown)+" seconds.</p>")

			port := os.Getenv("SMTP_PORT")
			portInt, _ := strconv.Atoi(port)
			d := gomail.NewDialer(os.Getenv("SMTP_HOST"),
				portInt,
				os.Getenv("SMTP_USER"),
				os.Getenv("SMTP_PASSWORD"))
			if err := d.DialAndSend(m); err != nil {
				s.Logger.Error(err)
			}
		}()

		w.Write([]byte("done!"))
	}
}
