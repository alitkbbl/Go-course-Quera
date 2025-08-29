package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type responseData struct {
	Status string                       `json:"status"`
	Stats  map[string]map[string]string `json:"stats"`
}

func GetExchangeRate(source, destination string) (string, error) {
	if source == "" {
		return "", errors.New("source currency is empty")
	}

	src := strings.ToLower(source)
	dst := strings.ToLower(destination)
	if dst == "" {
		dst = "rls"
	}

	endpoint := "http://localhost:4001/rates"
	params := url.Values{}
	params.Set("srcCurrency", src)
	params.Set("dstCurrency", dst)

	fullURL := endpoint + "?" + params.Encode()

	resp, err := http.Get(fullURL)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("server error %d: %s", resp.StatusCode, string(bodyBytes))
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	var data responseData
	err = json.Unmarshal(bodyBytes, &data)
	if err != nil {
		return "", fmt.Errorf("failed to parse JSON response: %w", err)
	}

	if data.Status != "OK" {
		return "", errors.New("API status not OK")
	}

	key := src + "-" + dst
	if val, exists := data.Stats[key]; exists {
		latest, ok := val["latest"]
		if !ok {
			return "", errors.New("latest price not found in response")
		}
		return latest, nil
	} else {
		return "", errors.New("price data not found for given currency pair")
	}
}
