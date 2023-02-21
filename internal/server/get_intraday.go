package server

import (
	"encoding/json"
	"net/http"

	"github.com/gpsinsight/go-interview-challenge/pkg/messages"
	"github.com/sirupsen/logrus"
)

type GetIntradayHandler func(w http.ResponseWriter, r *http.Request)

type IntradayListResponse struct {
	Items      []messages.IntradayValue `json:"items"`
	NextPageID int                      `json:"next_page_id,omitempty"`
}

func NewGetIntradayHandler(s messages.IntradayValueStore, l *logrus.Entry) GetIntradayHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		pageID := r.Context().Value(PageIDKey)
		page := pageID.(int)
		limit := 20 // hard set the number of items in a response
		offset := 0
		if page > 1 {
			offset = limit * (page - 1)
		}

		list, err := s.List(r.Context(), limit, offset)
		if err != nil {
			l.Errorf("error getting list: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		res := IntradayListResponse{
			Items:      list,
			NextPageID: pageID.(int) + 1,
		}

		b, err := json.Marshal(res)
		if err != nil {
			l.Errorf("error marshaling JSON: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("content-type", "application/json; charset=utf-8")
		_, err = w.Write(b)
		if err != nil {
			l.Errorf("error writing response: %s", err.Error())
			return
		}
		return
	}
}
