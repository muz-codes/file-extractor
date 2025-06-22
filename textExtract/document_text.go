package textExtract

import (
	"bytes"
	"encoding/json"
	"errors"
	"file-extractor/config"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

const modelId = "prebuilt-read"

var (
	key      = config.GetConfig("azure.ai_service_key")
	endpoint = config.GetConfig("azure.ai_service_endpoint")
	location = config.GetConfig("azure.ai_service_region")
)

type AnalyzedResult struct {
	Status          string `json:"status"`
	CreatedDateTime string `json:"createdDateTime"`
	LastUpdatedTime string `json:"lastUpdatedTime"`
	AnalyzeResult   struct {
		Content string `json:"content"`
	} `json:"analyzeResult"`
}

func ExtractDocument(file, docType string) (string, error) {
	// process file azure document OCR
	// Accepts url or base64 of file/docs
	// Create JSON body
	//postBody := &struct{ UrlSource string }{
	//	UrlSource: file,
	//}

	postBody := make(map[string]string)
	if docType == "base64" {
		postBody = map[string]string{"base64Source": file}
	} else if docType == "url" {
		postBody = map[string]string{"urlSource": file}
	}

	jsonBody, err := json.Marshal(postBody)
	if err != nil {
		log.Output(1, fmt.Sprintf("error in ExtractDocument: %v", err.Error()))
		return "", err
	}
	resultId, err := AnalyzeDocument(modelId, jsonBody)
	if err != nil {
		log.Output(1, fmt.Sprintf("error in ExtractDocument: %v", err.Error()))
		return "", err
	}
	analyzedResult, err := GetAnalyzeResult(modelId, resultId)
	if err != nil {
		log.Output(1, fmt.Sprintf("error in ExtractDocument: %v", err.Error()))
		return "", err
	}
	// need to hit another api to check for analysis
	return analyzedResult.AnalyzeResult.Content, nil
}

func AnalyzeDocument(modelId string, jsonBody []byte) (string, error) {
	// Define the URL
	url := fmt.Sprintf("%s/formrecognizer/documentModels/%s:analyze?api-version=2023-07-31&stringIndexType=textElements", endpoint, modelId)
	// Create new POST request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", err
	}

	// Set headers
	req.Host = strings.Trim(strings.Trim(strings.Trim(endpoint, "https://"), "http://"), "/")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Ocp-Apim-Subscription-Key", key)

	// Execute the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if !(200 <= resp.StatusCode && resp.StatusCode <= 299) {
		errorOccurred := fmt.Errorf("analyzeDocument response is %v", resp.StatusCode)
		return "", errorOccurred
	}

	// Read and print the response body
	_, err = io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	resultId := resp.Header.Get("apim-request-id")
	return resultId, nil
}

func GetAnalyzeResult(modelId string, resultId string) (AnalyzedResult, error) {
	var analyzedResult AnalyzedResult
	// Define the URL and replace placeholders with actual values
	url := fmt.Sprintf("%s/formrecognizer/documentModels/%s/analyzeResults/%s?api-version=2023-07-31", endpoint, modelId, resultId)

	// Create a new GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Output(1, fmt.Sprintf("error in GetAnalyzeResult: %v", err.Error()))
		return AnalyzedResult{}, err
	}

	// Set the Host header
	req.Host = strings.Trim(strings.Trim(strings.Trim(endpoint, "https://"), "http://"), "/")
	req.Header.Set("Ocp-Apim-Subscription-Key", key)

	// Execute the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Output(1, fmt.Sprintf("error in GetAnalyzeResult: %v", err.Error()))
		return AnalyzedResult{}, err
	}
	defer resp.Body.Close()
	if !(200 <= resp.StatusCode && resp.StatusCode <= 299) {
		errorOccurred := fmt.Errorf("analyzeDocument response is %v", resp.StatusCode)
		log.Output(1, fmt.Sprintf("error in GetAnalyzeResult: %v", errorOccurred.Error()))
		return AnalyzedResult{}, errorOccurred
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Output(1, fmt.Sprintf("error in GetAnalyzeResult: %v", err.Error()))
		return AnalyzedResult{}, err
	}

	// Print the response body
	err = json.Unmarshal(body, &analyzedResult)
	if err != nil {
		log.Output(1, fmt.Sprintf("error in GetAnalyzeResult: %v", err.Error()))
		return AnalyzedResult{}, err
	}
	if analyzedResult.Status == "succeeded" {
		return analyzedResult, nil
	} else if analyzedResult.Status == "failed" {
		errorOccurred := errors.New("get analyze result of image failed")
		log.Output(1, "error in GetAnalyzeResult: "+errorOccurred.Error())
		return AnalyzedResult{}, errorOccurred
	}
	time.Sleep(1 * time.Second)
	analyzedResult, err = GetAnalyzeResult(modelId, resultId)
	if err != nil {
		log.Output(1, fmt.Sprintf("error in GetAnalyzeResult: %v", err.Error()))
		return AnalyzedResult{}, err
	}
	return analyzedResult, nil
}
