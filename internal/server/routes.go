package server

import "net/http"

func (s *Server) BuildRoutes() {
	s.router.HandleFunc("/stocks/{ticker}", s.GetIntradayValues).Methods(http.MethodGet).Name("GetIntradayValues")
}
