package task

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/euclia/goJaqpot/models"
)

const (
	taskPath = "jaqpot/services/task/"
)

// GetTask is a method to get a Task by ID.
func GetTask(taskID string, AuthToken string, BaseURL string, HTTPClient *http.Client) (retTask models.Task, err error) {
	var endpoint = BaseURL + taskPath + taskID

	req, err := http.NewRequest("GET", endpoint, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+AuthToken)

	resp, err := HTTPClient.Do(req)

	var returnTask models.Task

	if err != nil {
		fmt.Printf(err.Error())
		return returnTask, err
	}

	if resp.StatusCode >= 300 {
		var errorReport models.ErrorReport
		_ = json.NewDecoder(resp.Body).Decode(&errorReport)
		defer resp.Body.Close()
		err = errors.New(errorReport.Message)
		return returnTask, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&returnTask)
	return returnTask, err
}
