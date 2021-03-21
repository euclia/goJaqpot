package feature

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/euclia/gojaqpot/models"
)

const (
	featurePath = "jaqpot/services/feature/"
)

// GetFeature is a method to get a feature by ID.
func GetFeature(featureID string, AuthToken string, BaseURL string, HTTPClient *http.Client) (feat models.Feature, err error) {
	var endpoint = BaseURL + featurePath + featureID
	req, err := http.NewRequest("GET", endpoint, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+AuthToken)

	resp, err := HTTPClient.Do(req)
	var returnFeat models.Feature

	if err != nil {
		fmt.Printf(err.Error())
		return returnFeat, err
	}

	if resp.StatusCode != 200 {
		var errorReport models.ErrorReport
		_ = json.NewDecoder(resp.Body).Decode(&errorReport)
		defer resp.Body.Close()
		err = errors.New(errorReport.Message)
		return returnFeat, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&returnFeat)
	return returnFeat, err
}
