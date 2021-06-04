package server

import "net/http"

func (s *Server) handleError(w http.ResponseWriter, message string) {
	s.logger.Error(message)
	w.Write([]byte("Sorry, there was an error! E: " + message))
}
