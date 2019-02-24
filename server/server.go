package server

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Server struct {
	Router *httprouter.Router
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	s.Router.ServeHTTP(w, r)
}
