package cve

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CVEStatus struct {
	Bugzilla       Bugzilla `json:"bugzilla"`
	ThreatSeverity string   `json:"threat_severity"`
}

type Bugzilla struct {
	Description string `json:"description"`
	ID          string `json:"id"`
	URL         string `json:"url"`
}

func getInfo(cveNumber string) (*CVEStatus, error) {
	url := fmt.Sprintf("https://access.redhat.com/labs/securitydataapi/cve/%s.json", cveNumber)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	cveStatus := CVEStatus{}
	err = json.NewDecoder(resp.Body).Decode(&cveStatus)
	if err != nil {
		return nil, err
	}
	return &cveStatus, nil
}
