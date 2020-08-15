package http

import (
	"flag"
	"fmt"
	"net/http"
)

type HttpServer struct {
	Host            string
	HostFlagKey     string
	HostFlagDefault string

	Port            string
	PortFlagKey     string
	PortFlagDefault string
}

func NewHttpServer() *HttpServer {
	return &HttpServer{
		Host:            "",
		HostFlagKey:     "host",
		HostFlagDefault: "localhost",
		Port:            "",
		PortFlagKey:     "port",
		PortFlagDefault: "8280",
	}
}

func (s *HttpServer) Address() string {
	return fmt.Sprintf("%s:%s", s.Host, s.Port)
}

func (s *HttpServer) StartMessage() string {
	return fmt.Sprintf("Starting HTTP server at %s:%s", s.Host, s.Port)
}

func (s *HttpServer) SetFlags() {
	flag.StringVar(&s.Host, s.HostFlagKey, s.HostFlagDefault, "host to run on")
	flag.StringVar(&s.Port, s.PortFlagKey, s.PortFlagDefault, "port to run on")
}

func (s *HttpServer) ListenAndServe(handler http.Handler) error {
	return http.ListenAndServe(s.Address(), handler)
}
