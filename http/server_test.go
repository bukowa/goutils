package http

import (
	"fmt"
	"testing"
)

func TestHttpServer_Address(t *testing.T) {
	type fields struct {
		Host string
		Port string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "1",
			fields: fields{
				Host: "0.0.0.0",
				Port: "80",
			},
		},
		{
			name: "2",
			fields: fields{
				Host: "1.1.1",
				Port: "11",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewServer()
			s.Host = tt.fields.Host
			s.Port = tt.fields.Port
			if s.Address() != fmt.Sprintf("%s:%s", tt.fields.Host, tt.fields.Port) {
				t.Error(tt)
			}
		})
	}
}

func TestHttpServer_StartMessage(t *testing.T) {
	var expected = "Starting HTTP server at xas:2999"
	s := NewServer()
	s.Host = "xas"
	s.Port = "2999"
	if msg := s.StartMessage(); msg != expected {
		t.Error(s.StartMessage())
	}
}
