package server

import "net/http"

func (s *Server) handleError(w http.ResponseWriter, message string) {
	w.Write([]byte("Sorry, there was an error! E: "+message))
}
