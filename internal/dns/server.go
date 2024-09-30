package dns

import (
	"fmt"
	"github.com/miekg/dns"
	"github.com/spf13/viper"
	"log"
	"math/rand"
	"strings"
	"time"
)

type Server struct {
	addr            string
	blocklist       map[string]struct{}
	upstreamServers []string
}

const seedValue = 1 // Constant for seed value if needed

// NewServer creates a new DNS server instance
func NewServer() *Server {
	// Read interface and port from the config
	interfaceAddr := viper.GetString("server.interface")
	port := viper.GetInt("server.port")

	addr := fmt.Sprintf("%s:%d", interfaceAddr, port)

	return &Server{
		addr:            addr,
		blocklist:       loadBlocklist(),       // Load blocklist from a separate config file
		upstreamServers: loadUpstreamServers(), // Load upstream servers from separate config file
	}
}

// Start the DNS server
func (s *Server) Start() error {
	dns.HandleFunc(".", s.handleDNSRequest)

	server := &dns.Server{Addr: s.addr, Net: "udp"}
	log.Printf("Starting DNS server on %s", s.addr)
	return server.ListenAndServe()
}

// handleDNSRequest processes incoming DNS queries
func (s *Server) handleDNSRequest(w dns.ResponseWriter, r *dns.Msg) {
	domain := strings.ToLower(r.Question[0].Name)

	// Check if the domain is in the blocklist
	if s.isBlocked(domain) {
		s.returnNegativeResponse(w, r)
		return
	}

	// Forward the request to an upstream DNS server
	upstream := s.getRandomUpstreamServer()
	m, err := dns.Exchange(r, upstream)
	if err != nil {
		log.Printf("Failed to resolve %s via %s: %v", domain, upstream, err)
		s.returnNegativeResponse(w, r)
		return
	}

	_ = w.WriteMsg(m)
}

// isBlocked checks if the domain is in the blocklist
func (s *Server) isBlocked(domain string) bool {
	_, exists := s.blocklist[domain]
	return exists
}

// returnNegativeResponse returns an NXDOMAIN (negative) response
func (s *Server) returnNegativeResponse(w dns.ResponseWriter, r *dns.Msg) {
	m := new(dns.Msg)
	m.SetRcode(r, dns.RcodeNameError) // NXDOMAIN
	_ = w.WriteMsg(m)
}

// loadBlocklist loads the blocklist from a separate config file
func loadBlocklist() map[string]struct{} {
	viper.SetConfigName("blocklist_config") // Name of the config file (without extension)
	viper.SetConfigType("yaml")             // Config file type (YAML)
	viper.AddConfigPath("./config/")        // Path to look for the config file

	err := viper.ReadInConfig() // Read the config file
	if err != nil {
		log.Fatalf("Error reading blocklist config file: %v", err)
	}

	blocklist := make(map[string]struct{})
	blockedDomains := viper.GetStringSlice("blocklist")
	if len(blockedDomains) == 0 {
		log.Printf("No domains found in the blocklist config file.")
	}

	for _, domain := range blockedDomains {
		// Ensure domains end with a dot for valid DNS comparison
		if !strings.HasSuffix(domain, ".") {
			domain += "."
		}
		blocklist[domain] = struct{}{}
	}
	return blocklist
}

// loadUpstreamServers loads upstream DNS servers from a separate config file
func loadUpstreamServers() []string {
	viper.SetConfigName("upstream_config") // Name of the config file (without extension)
	viper.SetConfigType("yaml")            // Config file type (YAML)
	viper.AddConfigPath("./config/")       // Path to look for the config file

	err := viper.ReadInConfig() // Read the config file
	if err != nil {
		log.Fatalf("Error reading upstream config file: %v", err)
	}

	upstreamServers := viper.GetStringSlice("upstream_servers")
	if len(upstreamServers) == 0 {
		log.Fatalf("No upstream DNS servers found in the configuration file.")
	}

	log.Printf("Loaded upstream DNS servers: %v", upstreamServers)
	return upstreamServers
}

// getRandomUpstreamServer returns a random upstream DNS server from the list
func (s *Server) getRandomUpstreamServer() string {
	// Create a new random generator with a local seed (based on time)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return s.upstreamServers[r.Intn(len(s.upstreamServers))]
}
