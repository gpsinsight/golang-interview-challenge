package server

import "github.com/gpsinsight/go-interview-challenge/pkg/messages"

func (server *Server) Route(s messages.IntradayValueStore) {
	/**
	 * TODO: Build endpoint for exposing intraday data
	 */
	server.router.HandleFunc("/api/v1/intraday", NewGetIntradayHandler(s, server.logger))
	server.router.Use(Pagination)
}
