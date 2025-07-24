package discovery

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type discoverResp struct {
	Target string `json:"target"`
}

// GetServiceAddress contacts the discovery-service and returns base URL of given service.
func GetServiceAddress(discoveryBaseURL, serviceName string) (string, error) {
	url := fmt.Sprintf("%s/discover/%s", discoveryBaseURL, serviceName)

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("service %s not found, status %d", serviceName, resp.StatusCode)
	}

	var dr discoverResp
	if err := json.NewDecoder(resp.Body).Decode(&dr); err != nil {
		return "", err
	}
	return dr.Target, nil
}
