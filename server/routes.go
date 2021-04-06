package server

import "net/http"

// here are all the routes and their handlers
func (s *Server) setRoutes() {
	// basic routes
	routes := map[string]http.HandlerFunc{
		"/":        s.handleIndex(),
		"/intent":  s.handleIntent(),
		"/receipt": s.handleReceipt(),
	}

	// wrap the routes in basic middleware stack
	for route, handler := range routes {
		s.router.HandleFunc(route, s.logRequest(handler))
	}

	// and finally, register the static assets that are bundled by rice
	assetServer := http.StripPrefix("/static/", http.FileServer(s.staticAssets.HTTPBox()))
	s.router.PathPrefix("/").Handler(assetServer)
}
