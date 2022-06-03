package main

import "net/http"

type Server struct {
	port   string
	router *Router
}

func NewServer(port string) *Server {
	return &Server{
		port:   port,
		router: NewRouter(),
	}
}

func (s *Server) Listen() error {
	http.Handle("/", s.router)
	err := http.ListenAndServe(s.port, nil)
	if err != nil {
		return err
	}
	return nil
}

func (server *Server) Handle(method string, path string, handler http.HandlerFunc) {
	_, exist := server.router.rules[path]
	if !exist {
		server.router.rules[path] = make(map[string]http.HandlerFunc)
	}
	server.router.rules[path][method] = handler
}

//Con esto se maneja los middleware
func (s *Server) AddMiddleware(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	//Usando los 3 puntos le digo que no se cuantos van a venir
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}
