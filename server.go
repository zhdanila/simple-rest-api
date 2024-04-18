package todo_list

import (
	"fmt"
	"net/http"
)

type Server struct {
	srv *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.srv = &http.Server{
		Addr:    ":" + port,
		Handler: handler,
	}

	fmt.Printf("server started on %s port\n", port)

	return s.srv.ListenAndServe()
}
