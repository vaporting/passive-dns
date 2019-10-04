package hunter

import (
	"passive-dns/db"

	"passive-dns/types"

	"passive-dns/models"

	"encoding/json"
)

// SourceIPHunter hunt targets with ips
type SourceIPHunter struct {
	*hunter
}

// Hunt hunts targets from sources
func (hunter *SourceIPHunter) Hunt(sources []string) ([]byte, error) {
	results := make(map[string][]types.TargetDomain)
	for _, source := range sources {
		result, _ := hunter.huntTargets(source)
		results[source] = result
	}

	jsonBytes, err := json.Marshal(results)
	return jsonBytes, err
}

// huntTarget hunts targets from source
func (hunter *SourceIPHunter) huntTargets(source string) ([]types.TargetDomain, error) {
	rEntries := []models.ResolvedIP{}
	err := hunter.db.Model(&rEntries).
		ColumnExpr("resolved_ip.first_seen, resolved_ip.last_seen").
		ColumnExpr("domain.name AS dname").
		Join(hunter.joinFmtCmd, source).
		Order("first_seen ASC").
		Select()

	results := make([]types.TargetDomain, len(rEntries))
	for index, entry := range rEntries {
		results[index] = make(map[string]*types.SeenGroup)
		results[index][entry.Dname] = &types.SeenGroup{
			FirstSeen: entry.FirstSeen.Format("2006-01-02"),
			LastSeen:  entry.LastSeen.Format("2006-01-02")}
	}
	return results, err
}

// NewSourceIPHunter creates SourceIPHunter
func NewSourceIPHunter() *SourceIPHunter {
	hunter := SourceIPHunter{hunter: &hunter{}}
	tempDB, _ := db.GetDB()
	hunter.db = tempDB
	hunter.SourceTypes = []string{types.SourceIpsType}
	hunter.joinFmtCmd =
		"INNER JOIN sources ON sources.id = resolved_ip.source_id" +
			" INNER JOIN domains AS domain ON domain.id = resolved_ip.domain_id" +
			" INNER JOIN ips ON ips.id = resolved_ip.resolved_ip_id" +
			" AND ips.ip = ?"
	return &hunter
}
