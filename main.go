package main

import (
	"log"
	"net/http"
)

type server struct {
	addr string
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		switch r.URL.Path {
		case "/":
			w.Write([]byte("Index Page"))
			return
		case "/users":
			w.Write([]byte("Users Page"))
			return
		}
	default:
		w.Write([]byte("Method not allowed"))
		return
	}
}

func main() {
	s := &server{addr: ":8082"}
	if err := http.ListenAndServe(s.addr, s); err != nil {
		log.Fatal(err)
	}
}
