package dns

import (
	"github.com/miekg/dns"
	"sync"
	"time"
)

// CacheEntry represents a cached DNS response with an expiration time
type CacheEntry struct {
	Msg        *dns.Msg
	Expiration time.Time
}

// DNSCache struct for the in-memory cache, with a mutex for thread safety
type DNSCache struct {
	data map[string]CacheEntry
	ttl  time.Duration
	mu   sync.RWMutex // RWMutex to handle concurrent access
}

// NewDNSCache initializes the cache with a given TTL
func NewDNSCache(ttl time.Duration) *DNSCache {
	return &DNSCache{
		data: make(map[string]CacheEntry),
		ttl:  ttl,
	}
}

// Get retrieves a DNS response from the cache if it exists and hasn't expired
// Also updates the ID field of the cached response to match the incoming request
func (c *DNSCache) Get(domain string, queryID uint16) (*dns.Msg, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	entry, found := c.data[domain]
	if !found || time.Now().After(entry.Expiration) {
		if found {
			delete(c.data, domain) // Remove expired entry
		}
		return nil, false
	}

	// Make a copy of the cached response and update the ID field to match the current query ID
	cachedMsg := entry.Msg.Copy()
	cachedMsg.Id = queryID
	return cachedMsg, true
}

// Set stores a DNS response in the cache with a TTL
func (c *DNSCache) Set(domain string, msg *dns.Msg) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[domain] = CacheEntry{
		Msg:        msg.Copy(),            // Store a copy of the message
		Expiration: time.Now().Add(c.ttl), // Set expiration time
	}
}
