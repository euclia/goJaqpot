package doa

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/euclia/goJaqpot/models"
)

const (
	doaPath = "jaqpot/services/doa/"
)

// GetDOA is a method to get a a model's DOA, by its ID.
func GetDOA(modelID string, AuthToken string, BaseURL string, HTTPClient *http.Client) (retDoa models.Doa, err error) {
	var endpoint = BaseURL + doaPath

	req, err := http.NewRequest("GET", endpoint, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+AuthToken)
	// req.Header.Set("Accept", "application/json")

	q := req.URL.Query()
	q.Add("hasSources", modelID)

	// req.URL.RawQuery = q.Encode()
	resp, err := HTTPClient.Do(req)
	var returnDoa models.Doa

	if err != nil {
		fmt.Printf(err.Error())
		return returnDoa, err
	}

	if resp.StatusCode != 200 {
		var errorReport models.ErrorReport
		_ = json.NewDecoder(resp.Body).Decode(&errorReport)
		defer resp.Body.Close()
		err = errors.New(errorReport.Message)
		return returnDoa, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&returnDoa)
	return returnDoa, err
}
