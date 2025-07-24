package consul

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

// RegisterToConsul registers the service instance to local Consul agent.
// serviceID must be unique (e.g. gateway-4000), serviceName shared (e.g. gateway-service).
func RegisterToConsul(serviceID, serviceName, serviceAddress string, servicePort int) {
	serviceAddress = "host.docker.internal"
	payload := fmt.Sprintf(`{
		"ID": "%s",
		"Name": "%s",
		"Address": "%s",
		"Port": %d,
		"Check": {
			"HTTP": "http://%s:%d/health",
			"Interval": "10s",
			"Timeout": "2s"
		}
	}`, serviceID, serviceName, serviceAddress, servicePort, serviceAddress, servicePort)

	req, err := http.NewRequest("PUT", "http://localhost:8500/v1/agent/service/register", bytes.NewBufferString(payload))
	if err != nil {
		log.Fatalf("❌ failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("❌ failed to register to Consul: %v", err)
	}
	defer resp.Body.Close()

	log.Println("✅ Consul response status:", resp.Status)
}
