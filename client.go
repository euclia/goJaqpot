package gojaqpot

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/euclia/gojaqpot/dataset"
	"github.com/euclia/gojaqpot/doa"
	"github.com/euclia/gojaqpot/feature"
	"github.com/euclia/gojaqpot/model"
	"github.com/euclia/gojaqpot/models"
	"github.com/euclia/gojaqpot/task"
)

const (
	// ClientVersion is used in User-Agent request header to provide server with API level.
	ClientVersion = "0.0.1"

	// httpClientTimeout is used to limit http.Client waiting time.
	httpClientTimeout = 15 * time.Second
)

// Client object
type Client models.Client

// InitClient creates a Jaqpot Go Client
func InitClient(baseURL string) *Client {
	if baseURL[len(baseURL)-1:] != "/" {
		baseURL = baseURL + "/"
	}
	return &Client{
		C: models.ClientProperties{
			BaseURL: baseURL,
			HTTPClient: &http.Client{
				Timeout: 100000000000,
			},
		},
	}
}

// IJaqpotClient is the Jaqpot Client Interface
type IJaqpotClient interface {
	// Get a Jaqpot Feature by its id.
	GetFeature(featureID string, AuthToken string) (feat models.Feature, err error)

	// Get a Jaqpot Dataset by its id.
	GetDataset(datasetID string, AuthToken string) (data models.Dataset, err error)

	// Get a model's DOA by the model's id.
	GetDOA(modelID string, AuthToken string) (modelDoa models.Doa, err error)

	// Get a Task by its id.
	GetTask(taskID string, AuthToken string) (returnTask models.Task, err error)

	// GetModel is a method to get a model by ID.
	GetModel(modelID string, AuthToken string) (retModel models.Model, err error)

	// GetMyModels is a method to get a list of user's models.
	GetMyModels(min int, max int, AuthToken string) (myModels models.Models, err error)

	// GetOrgsModels is a method to get a list of an organization's models.
	GetOrgsModels(organizationID string, min int, max int, AuthToken string) (orgsModels models.Models, err error)

	// GetOrgsModelsByTag is a method to get a list of an organization's models with a particular tag.
	GetOrgsModelsByTag(organizationID string, tag string, min int, max int, AuthToken string) (tagModels models.Models, err error)

	// Predict is a method to make a prediction on a Jaqpot Dataset (returns the task ID).
	Predict(modelID string, values []map[string]interface{}, AuthToken string) (prediction models.Prediction, err error)
}

// GetFeature is a method to get a feature by ID.
func (client *Client) GetFeature(featureID string, AuthToken string) (feat models.Feature, err error) {
	return feature.GetFeature(featureID, AuthToken, client.C.BaseURL, client.C.HTTPClient)
}

// GetDataset is a method to get a Dataset by ID.
func (client *Client) GetDataset(datasetID string, AuthToken string) (data models.Dataset, err error) {
	return dataset.GetDataset(datasetID, AuthToken, client.C.BaseURL, client.C.HTTPClient)
}

// GetDOA is a method to get a model's DOA, by its ID.
func (client *Client) GetDOA(modelID string, AuthToken string) (modelDoa models.Doa, err error) {
	return doa.GetDOA(modelID, AuthToken, client.C.BaseURL, client.C.HTTPClient)
}

// GetTask is a method to get a Task its ID.
func (client *Client) GetTask(taskID string, AuthToken string) (returnTask models.Task, err error) {
	return task.GetTask(taskID, AuthToken, client.C.BaseURL, client.C.HTTPClient)
}

// GetModel is a method to get a model by ID.
func (client *Client) GetModel(modelID string, AuthToken string) (retModel models.Model, err error) {
	return model.GetModel(modelID, AuthToken, client.C.BaseURL, client.C.HTTPClient)
}

// GetMyModels is a method to get a list of user's models.
func (client *Client) GetMyModels(min int, max int, AuthToken string) (myModels models.Models, err error) {
	return model.GetMyModels(min, max, AuthToken, client.C.BaseURL, client.C.HTTPClient)
}

// GetOrgsModels is a method to get a list of an organization's models.
func (client *Client) GetOrgsModels(organizationID string, min int, max int, AuthToken string) (orgsModels models.Models, err error) {
	return model.GetOrgsModels(organizationID, min, max, AuthToken, client.C.BaseURL, client.C.HTTPClient)
}

// GetOrgsModelsByTag is a method to get a list of an organization's models with a particular tag.
func (client *Client) GetOrgsModelsByTag(organizationID string, tag string, min int, max int, AuthToken string) (tagModels models.Models, err error) {
	return model.GetOrgsModelsByTag(organizationID, tag, min, max, AuthToken, client.C.BaseURL, client.C.HTTPClient)
}

// Predict is a method to make a prediction on a Jaqpot Dataset (returns the task ID).
func (client *Client) Predict(modelID string, values []map[string]interface{}, AuthToken string) (prediction models.Prediction, err error) {

	var retPrediction models.Prediction
	var predTask models.Task

	jaqDataset := dataset.CreateDataset(modelID, values, AuthToken, client.C.BaseURL, client.C.HTTPClient)

	datasetID, internalError := dataset.PostDataset(jaqDataset, AuthToken, client.C.BaseURL, client.C.HTTPClient)

	taskID, internalError := model.Predict(modelID, datasetID, AuthToken, client.C.BaseURL, client.C.HTTPClient)

	if internalError != nil {
		fmt.Printf(internalError.Error())
		return retPrediction, internalError
	}

	for true {
		predTask, internalError = task.GetTask(taskID.SlashID, AuthToken, client.C.BaseURL, client.C.HTTPClient)
		percent := predTask.PercentageCompleted
		dataID := predTask.Result

		if percent == 100 {
			retPrediction.ModelID = modelID
			retPrediction.DatasetID = strings.Split(dataID, "/")[1]
			break
		} else {
			time.Sleep(1 * time.Second)
		}

	}

	if internalError != nil {
		fmt.Printf(internalError.Error())
		return retPrediction, internalError
	}

	retPrediction.Data, retPrediction.Predictions, internalError = formatPreds(retPrediction.DatasetID, AuthToken, client.C.BaseURL, client.C.HTTPClient)

	if internalError != nil {
		fmt.Printf(internalError.Error())
		return retPrediction, internalError
	}

	return retPrediction, err
}

func formatPreds(datasetID string, AuthToken string, BaseURL string, HTTPClient *http.Client) (data []map[string]interface{}, preds []map[string]interface{}, err error) {
	predDataset, internalError := dataset.GetDataset(datasetID, AuthToken, BaseURL, HTTPClient)
	var endpoint string
	reverse := make([]string, len(predDataset.Features))
	var retData []map[string]interface{}
	var retPreds []map[string]interface{}

	if internalError != nil {
		fmt.Printf(internalError.Error())
		return retData, retPreds, err
	}

	for _, item := range predDataset.Features {
		if item.Category == "PREDICTED" {
			endpoint = item.Name
		}
		myKey, _ := strconv.Atoi(item.Key)
		reverse[myKey] = item.Name
	}

	for _, item := range predDataset.DataEntry {

		currData := make(map[string]interface{})

		for key, val := range item.Values {
			currIndex, _ := strconv.Atoi(key)

			if reverse[currIndex] != endpoint {
				currData[reverse[currIndex]] = val
			} else {
				retPreds = append(retPreds, map[string]interface{}{endpoint: val})
			}
		}
		retData = append(retData, currData)
	}

	return retData, retPreds, err

}
