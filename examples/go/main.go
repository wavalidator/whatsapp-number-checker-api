// Check if a phone number has WhatsApp with the WA Validator API.
// Usage:
//
//	export WAVALIDATOR_API_KEY=your_key
//	go run main.go 14155551234
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

const apiURL = "https://wavalidator.com/api/v1/validate/"

type ValidateResponse struct {
	PhoneNumber      string `json:"phone_number"`
	Status           string `json:"status"`
	CreditsRemaining int    `json:"credits_remaining"`
	Cached           bool   `json:"cached,omitempty"`
}

func CheckWhatsAppNumber(phoneNumber, apiKey string) (*ValidateResponse, error) {
	body, _ := json.Marshal(map[string]string{"phone_number": phoneNumber})
	req, err := http.NewRequest(http.MethodPost, apiURL, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP %d", resp.StatusCode)
	}
	var result ValidateResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func main() {
	number := "14155551234"
	if len(os.Args) > 1 {
		number = os.Args[1]
	}

	result, err := CheckWhatsAppNumber(number, os.Getenv("WAVALIDATOR_API_KEY"))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch result.Status {
	case "valid":
		fmt.Printf("%s: has WhatsApp\n", result.PhoneNumber)
	case "invalid":
		fmt.Printf("%s: no WhatsApp\n", result.PhoneNumber)
	default:
		fmt.Printf("%s: check failed, try again later\n", result.PhoneNumber)
	}
	fmt.Printf("Credits remaining: %d\n", result.CreditsRemaining)
}
