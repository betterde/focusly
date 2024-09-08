package utils

import "net/http"

var camelToSnakeMapping = map[string]string{}

func CamelToSnake(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			q := r.URL.Query()
			for camel, snake := range camelToSnakeMapping {
				if q.Has(camel) {
					v := q.Get(camel)
					q.Del(camel)
					q.Set(snake, v)
				}
			}

			r.URL.RawQuery = q.Encode()
		}

		h.ServeHTTP(w, r)
	})
}
