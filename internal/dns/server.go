package dns

import (
	"github.com/miekg/dns"
	"log"
)

type Server struct {
	addr string
}

func NewServer() *Server {
	return &Server{
		addr: ":8455",
	}
}

func (s *Server) Start() error {
	dns.HandleFunc(".", handleDNSRequest)

	server := &dns.Server{Addr: s.addr, Net: "udp"}
	log.Printf("Starting DNS server on %s", s.addr)
	return server.ListenAndServe()
}

func handleDNSRequest(w dns.ResponseWriter, r *dns.Msg) {
	m := new(dns.Msg)
	m.SetReply(r)
	err := w.WriteMsg(m)
	if err != nil {
		return
	}
}
