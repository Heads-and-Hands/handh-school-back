package main

import (
	"github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"
	"handh-school-back/handlers"
	"net/http"
	"handh-school-back/myAdminConf"
	"strings"
	"encoding/base64"
	"handh-school-back/config"
)

func main() {
	println("Hello, school on 7771 port!")

	r := mux.NewRouter()
	r.Handle("/", handlers.GetHandler).Methods("GET")
	r.Handle("/", handlers.PostHandler).Methods("POST")

	m := myAdminConf.InitAdmin()
	r.PathPrefix("/admin").Handler(m)

	r.Use(Middleware)

	http.ListenAndServe(":7771", r)
}


func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !basicAuth(w, r) {
			w.Header().Set("WWW-Authenticate", `Basic realm="Beware! Protected REALM! "`)
			w.WriteHeader(401)
			w.Write([]byte("401 Unauthorized\n"))
		}
		next.ServeHTTP(w, r)
	})
}


func basicAuth(_ http.ResponseWriter, r *http.Request) bool {
	if r.URL.Path != "/admin" {
		return true
	}

	realPair := []string{config.User, config.Password}
	s := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
	if len(s) != 2 { return false }
	b, err := base64.StdEncoding.DecodeString(s[1])
	if err != nil {	return false }
	pair := strings.SplitN(string(b), ":", 2)
	if len(pair) != 2 {	return false }

	return pair[0] == realPair[0] && pair[1] == realPair[1]
}