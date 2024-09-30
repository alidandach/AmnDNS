package dns

import (
	"github.com/miekg/dns"
	"time"
)

type CacheEntry struct {
	msg        *dns.Msg
	expiration time.Time
}

type DNSCache struct {
	data map[string]CacheEntry
	ttl  time.Duration
}

// NewDNSCache initializes the cache with a specified TTL
func NewDNSCache(ttl time.Duration) *DNSCache {
	return &DNSCache{
		data: make(map[string]CacheEntry),
		ttl:  ttl,
	}
}

// Set adds a new entry to the cache
func (c *DNSCache) Set(domain string, msg *dns.Msg) {
	c.data[domain] = CacheEntry{
		msg:        msg,
		expiration: time.Now().Add(c.ttl),
	}
}

// Get retrieves an entry from the cache if it's still valid
func (c *DNSCache) Get(domain string) (*dns.Msg, bool) {
	entry, found := c.data[domain]
	if !found || time.Now().After(entry.expiration) {
		// If the entry is expired, delete it
		delete(c.data, domain)
		return nil, false
	}
	return entry.msg, true
}
