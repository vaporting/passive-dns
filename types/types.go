package types

// ResolvedEntry is used to store the resolved entry from request
type ResolvedEntry struct {
	Name      string `json:"name"`
	Type      string `json:"type"`
	Value     string `json:"value"`
	FirstSeen string `json:"first_seen"`
	LastSeen  string `json:"last_seen"`
}
