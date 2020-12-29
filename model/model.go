package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/euclia/goJaqpot/models"
)

const (
	modelPath = "jaqpot/services/model/"
)

// GetModel is a method to get a model by ID.
func GetModel(modelID string, AuthToken string, BaseURL string, HTTPClient *http.Client) (retModel models.Model, err error) {
	var endpoint = BaseURL + modelPath + modelID
	req, err := http.NewRequest("GET", endpoint, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+AuthToken)

	resp, err := HTTPClient.Do(req)
	var returnModel models.Model

	if err != nil {
		fmt.Printf(err.Error())
		return returnModel, err
	}

	if resp.StatusCode != 200 {
		var errorReport models.ErrorReport
		_ = json.NewDecoder(resp.Body).Decode(&errorReport)
		defer resp.Body.Close()
		err = errors.New(errorReport.Message)
		return returnModel, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&returnModel)
	return returnModel, err
}

// GetMyModels is a method to get a list of user's models.
func GetMyModels(min int, max int, AuthToken string, BaseURL string, HTTPClient *http.Client) (myModels models.Models, err error) {
	var endpoint = BaseURL + modelPath
	req, err := http.NewRequest("GET", endpoint, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+AuthToken)

	q := req.URL.Query()
	q.Add("min", strconv.Itoa(min))
	q.Add("max", strconv.Itoa(max))

	req.URL.RawQuery = q.Encode()
	resp, err := HTTPClient.Do(req)
	var returnModels models.Models

	if err != nil {
		fmt.Printf(err.Error())
		return returnModels, err
	}

	if resp.StatusCode != 200 {
		var errorReport models.ErrorReport
		_ = json.NewDecoder(resp.Body).Decode(&errorReport)
		defer resp.Body.Close()
		err = errors.New(errorReport.Message)
		return returnModels, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&returnModels.Models)

	returnModels.Total, _ = strconv.Atoi(resp.Header["Total"][0])

	return returnModels, err
}

// GetOrgsModels is a method to get a list of an organization's models.
func GetOrgsModels(organizationID string, min int, max int, AuthToken string, BaseURL string, HTTPClient *http.Client) (orgsModels models.Models, err error) {
	var endpoint = BaseURL + modelPath
	req, err := http.NewRequest("GET", endpoint, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+AuthToken)

	q := req.URL.Query()
	q.Add("organization", organizationID)
	q.Add("min", strconv.Itoa(min))
	q.Add("max", strconv.Itoa(max))

	req.URL.RawQuery = q.Encode()
	resp, err := HTTPClient.Do(req)
	var returnModels models.Models

	if err != nil {
		fmt.Printf(err.Error())
		return returnModels, err
	}

	if resp.StatusCode != 200 {
		var errorReport models.ErrorReport
		_ = json.NewDecoder(resp.Body).Decode(&errorReport)
		defer resp.Body.Close()
		err = errors.New(errorReport.Message)
		return returnModels, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&returnModels.Models)

	returnModels.Total, _ = strconv.Atoi(resp.Header["Total"][0])

	return returnModels, err
}

// GetOrgsModelsByTag is a method to get a list of an organization's models with a particular tag.
func GetOrgsModelsByTag(organizationID string, tag string, min int, max int, AuthToken string, BaseURL string, HTTPClient *http.Client) (tagModels models.Models, err error) {
	var endpoint = BaseURL + modelPath
	req, err := http.NewRequest("GET", endpoint, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+AuthToken)

	q := req.URL.Query()
	q.Add("organization", organizationID)
	q.Add("tag", tag)
	q.Add("min", strconv.Itoa(min))
	q.Add("max", strconv.Itoa(max))

	req.URL.RawQuery = q.Encode()
	resp, err := HTTPClient.Do(req)
	var returnModels models.Models

	if err != nil {
		fmt.Printf(err.Error())
		return returnModels, err
	}

	if resp.StatusCode != 200 {
		var errorReport models.ErrorReport
		_ = json.NewDecoder(resp.Body).Decode(&errorReport)
		defer resp.Body.Close()
		err = errors.New(errorReport.Message)
		return returnModels, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&returnModels.Models)

	returnModels.Total, _ = strconv.Atoi(resp.Header["Total"][0])

	return returnModels, err
}

// Predict is a method to make a prediction on a Jaqpot Dataset (returns the task ID).
func Predict(modelID string, datasetID string, AuthToken string, BaseURL string, HTTPClient *http.Client) (predictTask models.Task, err error) {
	var endpoint = BaseURL + modelPath + modelID

	body := url.Values{}
	body.Set("dataset_uri", BaseURL+"dataset/"+datasetID)
	body.Set("visible", "true")

	req, err := http.NewRequest("POST", endpoint, strings.NewReader(body.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Bearer "+AuthToken)

	resp, err := HTTPClient.Do(req)

	var returnTask models.Task

	if err != nil {
		fmt.Printf(err.Error())
		return returnTask, err
	}

	if resp.StatusCode != 200 {
		var errorReport models.ErrorReport
		_ = json.NewDecoder(resp.Body).Decode(&errorReport)
		defer resp.Body.Close()
		err = errors.New(errorReport.Message)
		return returnTask, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&returnTask)

	// var currList = strings.Split(resp.Header["Location"][0], "/")
	// returnID = currList[len(currList)-1]
	return returnTask, err
}
