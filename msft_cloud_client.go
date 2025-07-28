package main

import (
	// "encoding/json"
	// "regexp"

	"io"
	"net/http"
	"os"
	"time"
)

type MsftCloudClient struct {
	Logger GoLogger
	Client http.Client
	Field  string "message"
}

func NewMsftCloudClient(_logger GoLogger) *MsftCloudClient {
	client := http.Client{Timeout: 15 * time.Second}
	msftCloudClient := MsftCloudClient{Logger: _logger, Client: client, Field: ""}

	return &msftCloudClient
}

func (gpt *MsftCloudClient) GetResponse(w *http.ResponseWriter, r *http.Request) *http.Response {
	accessToken := os.Getenv("ONEDRIVE_ACCESS_TOKEN") // Set your token via env

	// req, err := http.NewRequest("GET", "https://", nil)

	// This URL gets the list of items in the root directory
	// You can customize it to target a specific folder or image by ID or path
	url := os.Getenv("JN_WEDDING_PHOTOS_PATH")

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		gpt.Logger.Error("Error creating request:" + err.Error())
		return &http.Response{StatusCode: req.Response.StatusCode, Status: "Error Request failed : " + req.Response.Status}
	}
	req.Header.Add("Authorization", "Bearer "+accessToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		gpt.Logger.Error("Request failed:" + err.Error())
		return &http.Response{StatusCode: resp.StatusCode, Status: "Error Request failed : " + resp.Request.Response.Status}
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		gpt.Logger.Error("Error reading response:" + err.Error())
		return &http.Response{StatusCode: resp.StatusCode, Status: "Error Request failed : " + resp.Request.Response.Status}
	}

	gpt.Logger.Info("Response:" + string(body))

	return resp
}
