package server

import (
	"html/template"
	"net/http"
)

func (s *Server) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tplString, err := s.views.String("index.tpl")
		if err != nil {
			s.handleError(w, err.Error())
			return
		}

		tpl, err := template.New("homepage").Parse(tplString)
		s.log(s.configuration)
		tpl.Execute(w, map[string]string{
			"Client":          s.configuration.customerName,
			"StripePublicKey": s.configuration.stripePublicKey,
		})
	}
}
