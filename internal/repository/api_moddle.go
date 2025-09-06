package repository

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func NewApiModel(endpoint string) *ApiModel {
	return &ApiModel{Endpoint: endpoint}
}

type ApiModel struct {
	Endpoint string
}

func (r *ApiModel) HashingPassword(password string) string {
	// Make a POST request to the endpoint to get the hashed password
	payload := map[string]string{"password": password}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return ""
	}

	req, err := http.NewRequest("POST", r.Endpoint+"/api_hashing.php", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return ""
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	// Print the equivalent curl command for debugging
	curlCmd := "curl -X POST " + r.Endpoint + "/api_hashing.php -H 'Content-Type: application/json' -d '" + string(jsonPayload) + "'"
	println("CURL:", curlCmd)

	var result struct {
		Hash string `json:"hash"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return ""
	}

	return result.Hash
}
