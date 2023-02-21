package server

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
)

const PageIDKey = "page"

func Pagination(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pageID := r.URL.Query().Get(PageIDKey)
		intPageID := 1
		var err error
		if pageID != "" {
			intPageID, err = strconv.Atoi(pageID)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				_, err := w.Write([]byte(fmt.Sprintf("invalid page id, %s must be an integer", PageIDKey)))
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
				}
			}
		}
		ctx := context.WithValue(r.Context(), PageIDKey, intPageID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
