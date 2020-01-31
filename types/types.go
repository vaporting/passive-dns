package types

import "time"

// Config structure
type Config struct {
	DB struct {
		Host string `yaml:"host" envconfig:"DB_HOST" required:"true"`
		Port string `yaml:"port" envconfig:"DB_PORT" required:"true"`
		Name string `yaml:"name" envconfig:"DB_NAME" required:"true"`
		User string `yaml:"user" envconfig:"DB_USER" required:"true"`
		PWD  string `yaml:"pwd" envconfig:"DB_PWD" required:"true"`
	} `yaml:"DB"`
	CACHE struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		PWD      string `yaml:"pwd"`
		PoolSize string `yaml:"pool_size"`
	} `yaml:"CACHE"`
	SyncSer struct {
		Writer string `yaml:"writer"`
	} `yaml:"Sync_service"`
}

// DNS record types
const (
	// DNSIpv4Type record type
	DNSIpv4Type string = "A"
	// DNSIpv6Type record type
	DNSIpv6Type string = "AAAA"
	// DNSDomainType record type
	DNSDomainType string = "CNAME"
)

// hunters support source types
const (
	// SourceIpsType source type
	SourceIpsType string = "ips"
	// SourceDomainsType source type
	SourceDomainsType string = "domains"
	// SourcePDomainsType source type
	SourcePDomainsType string = "parent_domains"
)

// ResolvedEntry is used to store the resolved entry from request
type ResolvedEntry struct {
	Name      string `json:"name"`
	Type      string `json:"type"`
	Value     string `json:"value"`
	FirstSeen string `json:"first_seen"`
	LastSeen  string `json:"last_seen"`
}

// ResolvedRow is the data before store to resolved table
type ResolvedRow struct {
	OriginalID uint
	PassiveID  uint
	SourceID   uint
	FirstSeen  time.Time
	LastSeen   time.Time
}

// HuntingSources uses to store all sources from request
type HuntingSources struct {
	Ips     []string
	Domains []string
	// ParentDomaind []string // currently, not using
	// StartDate time.Time // currently, not using
	// EndDate time.Time //currently, not using
}

// SeenGroup stores pair of first_seen and last_seen
type SeenGroup struct {
	FirstSeen string `json:"first_seen"`
	LastSeen  string `json:"last_seen"`
}

// TargetName is used to build result of hunting by source:ip
type TargetName map[string]*SeenGroup
