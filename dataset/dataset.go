package dataset

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/euclia/gojaqpot/model"
	"github.com/euclia/gojaqpot/models"
)

const (
	datasetPath = "jaqpot/services/dataset/"
)

// GetDataset is a method to get a Jaqpot Dataset by ID.
func GetDataset(datasetID string, AuthToken string, BaseURL string, HTTPClient *http.Client) (dataset models.Dataset, err error) {
	var endpoint = BaseURL + datasetPath + datasetID
	req, err := http.NewRequest("GET", endpoint, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+AuthToken)
	req.Header.Set("Accept", "application/json")

	q := req.URL.Query()
	q.Add("dataEntries", "true")

	req.URL.RawQuery = q.Encode()
	resp, err := HTTPClient.Do(req)
	var returnData models.Dataset

	if err != nil {
		fmt.Printf(err.Error())
		return returnData, err
	}

	if resp.StatusCode != 200 {
		var errorReport models.ErrorReport
		_ = json.NewDecoder(resp.Body).Decode(&errorReport)
		defer resp.Body.Close()
		err = errors.New(errorReport.Message)
		return returnData, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&returnData)
	return returnData, err
}

// PostDataset is a method to post a Jaqpot Dataset.
func PostDataset(data models.Dataset, AuthToken string, BaseURL string, HTTPClient *http.Client) (SlashID string, err error) {
	var endpoint = BaseURL + datasetPath
	body, err := json.Marshal(data)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+AuthToken)
	// req.Header.Set("Accept", "application/json")

	resp, err := HTTPClient.Do(req)
	// fmt.Println(resp.Body)
	var returnID string
	var returnData models.Dataset

	if err != nil {
		fmt.Printf(err.Error())
		return returnID, err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&returnData)

	var currList = strings.Split(resp.Header["Location"][0], "/")
	returnID = currList[len(currList)-1]
	return returnID, err
}

// CreateDataset is a method to create a Dataset object (used for the predict method).
func CreateDataset(modelID string, values []map[string]interface{}, AuthToken string, BaseURL string, HTTPClient *http.Client) (dataset models.Dataset) {
	var info models.FeatureInfo
	var returnData models.Dataset

	var cnt = 0
	reverse := make(map[string]string)
	currentModel, _ := model.GetModel(modelID, AuthToken, BaseURL, HTTPClient)

	var independentFeatures = currentModel.AdditionalInfo.(map[string]interface{})["independentFeatures"]

	for index, value := range independentFeatures.(map[string]interface{}) {

		// Dynamically add a sub-map
		info.URI = index
		info.Key = strconv.Itoa(cnt)
		info.Name = fmt.Sprintf("%v", value)
		reverse[info.Name] = strconv.Itoa(cnt)

		// The slice grows as needed.
		returnData.Features = append(returnData.Features, info)
		cnt++
	}

	cnt = 0

	var data models.DataEntry
	var entry models.EntryID
	vals := make(map[string]interface{})

	for _, item := range values {
		entry.Name = strconv.Itoa(cnt)

		for index, value := range item {
			vals[reverse[index]] = value
		}

		data.Values = vals
		data.EntryID = entry

		returnData.DataEntry = append(returnData.DataEntry, data)
		vals = make(map[string]interface{})
		cnt++

	}
	return returnData
}
