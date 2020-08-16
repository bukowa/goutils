package http

import (
	"flag"
	"fmt"
	"net/http"
)

type Server struct {
	FlagSet *flag.FlagSet

	Host            string
	HostFlagKey     string
	HostFlagDefault string

	Port            string
	PortFlagKey     string
	PortFlagDefault string
}

func NewServer() *Server {
	return &Server{
		FlagSet:         flag.CommandLine,
		Host:            "",
		HostFlagKey:     "host",
		HostFlagDefault: "localhost",
		Port:            "",
		PortFlagKey:     "port",
		PortFlagDefault: "8280",
	}
}

func (s *Server) Address() string {
	return fmt.Sprintf("%s:%s", s.Host, s.Port)
}

func (s *Server) StartMessage() string {
	return fmt.Sprintf("Starting HTTP server at %s:%s", s.Host, s.Port)
}

func (s *Server) SetFlags() {
	s.FlagSet.StringVar(&s.Host, s.HostFlagKey, s.HostFlagDefault, "host to run on")
	s.FlagSet.StringVar(&s.Port, s.PortFlagKey, s.PortFlagDefault, "port to run on")
}

func (s *Server) ListenAndServe(handler http.Handler) error {
	return http.ListenAndServe(s.Address(), handler)
}
