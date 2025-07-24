package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetServicesFromConsul(serviceName string) ([]ServiceInfo, error) {
	url := fmt.Sprintf("http://localhost:8500/v1/catalog/service/%s", serviceName)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []ServiceInfo
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}
