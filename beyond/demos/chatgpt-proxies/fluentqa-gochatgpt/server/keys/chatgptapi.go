package keys

// This file contains checks for the validity of OpenAI API keys and how much credit is left.

import (
	"encoding/json"
	"net/http"
)

func GetCredits(apiKey string) (BillingSubscription, error) {
	// Make request
	req, err := http.NewRequest("GET", "https://api.openai.com/dashboard/billing/subscription", nil)
	if err != nil {
		return BillingSubscription{}, err
	}
	req.Header.Add("Authorization", "Bearer "+apiKey)
	req.Header.Add("Content-Type", "application/json")
	// Send request
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return BillingSubscription{}, err
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return BillingSubscription{}, err
	}
	// Parse response
	var creditSummary BillingSubscription
	// Map response to struct
	err = json.NewDecoder(response.Body).Decode(&creditSummary)
	if err != nil {
		return BillingSubscription{}, err
	}
	return creditSummary, nil
}

func GetGrants(apiKey string) (CreditSummary, error) {
	req, err := http.NewRequest("GET", "https://api.openai.com/dashboard/billing/credit_grants", nil)
	if err != nil {
		return CreditSummary{}, err
	}
	req.Header.Add("Authorization", "Bearer "+apiKey)
	req.Header.Add("Content-Type", "application/json")
	// Send request
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return CreditSummary{}, err
	}
	defer response.Body.Close()
	var creditSummary CreditSummary
	// Map response to struct
	err = json.NewDecoder(response.Body).Decode(&creditSummary)
	if err != nil {
		return CreditSummary{}, err
	}
	if response.StatusCode != 200 {
		return creditSummary, err
	}
	return creditSummary, nil
}

func GetTotalCredits(apiKey string) (float64, error) {
	grants, err := GetGrants(apiKey)
	if err != nil {
		return 0, err
	}
	grant_remaining := grants.TotalAvailable
	subscription, err := GetCredits(apiKey)
	if err != nil {
		return 0, err
	}
	if subscription.HardLimitUSD < 19 {
		return grant_remaining, nil
	}
	return subscription.SoftLimitUSD, nil
}
