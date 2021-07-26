package main

import (
	"crypto/subtle"

	"github.com/Heads-and-Hands/handh-school-back/config"
	"github.com/Heads-and-Hands/handh-school-back/handlers"
	"github.com/Heads-and-Hands/handh-school-back/myAdminConf"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	_ "github.com/Heads-and-Hands/handh-school-back/bindatafs"

	"net/http"
)

func main() {
	println("Hello, school on 7771 port!")

	r := mux.NewRouter()

	m := myAdminConf.InitAdmin()
	r.PathPrefix("/admin").Handler(m)

	r.Handle("/", handlers.GetHandler).Methods("GET")
	r.Handle("/", handlers.PostHandler).Methods("POST")

	r.Use(AuthMiddleware)

	http.ListenAndServe(":7771", r)
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !basicAuth(w, r) {
			w.Header().Set("WWW-Authenticate", `Basic realm="Beware! Protected REALM! "`)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("401 Unauthorized\n"))
			return
		}
		next.ServeHTTP(w, r)
	})
}

func basicAuth(_ http.ResponseWriter, r *http.Request) bool {
	if r.URL.Path != "/admin" {
		return true
	}

	user, pwd, ok := r.BasicAuth()
	if !ok {
		return false
	}

	return subtle.ConstantTimeCompare([]byte(user), []byte(config.User)) == 1 &&
		subtle.ConstantTimeCompare([]byte(pwd), []byte(config.Password)) == 1
}
