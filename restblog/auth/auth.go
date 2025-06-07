package middleware

import "net/http"

func RequireAuth(f http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		token := r.Header.Get("Authorization")
		if token != "Bearer Secret" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		}
		f.ServeHTTP(w, r)
	})
}
