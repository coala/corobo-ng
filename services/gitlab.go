package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Get Gitlab access token from the given authorization code
func GetGitlabAccessToken(code string) (string, error) {
	var err error

	jsonBody, _ := json.Marshal(map[string]string{
		"client_id":     os.Getenv("GITLAB_CLIENT_ID"),
		"client_secret": os.Getenv("GITLAB_CLIENT_SECRET"),
		"code":          code,
		"grant_type":    "authorization_code",
		"redirect_uri":  "http://localhost:3000/login/gitlab/complete",
	})
	client := &http.Client{}
	req, _ := http.NewRequest("POST", "https://gitlab.com/oauth/token", bytes.NewBuffer(jsonBody))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	log.Printf("Requesting Gitlab for access token for code %s", code)
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error getting access token for Gitlab, %v", err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var tokenData map[string]interface{}
	err = json.Unmarshal(body, &tokenData)

	if err != nil {
		log.Printf("Error decoding Gitlab token response, %v", err)
		return "", err
	}

	if _, ok := tokenData["error"]; ok {
		log.Printf("Received an error response from Gitlab while requesting access token")
		return "", error(fmt.Errorf("error response, %v", tokenData))
	}

	log.Printf("Successfully retrieved access token, %s", tokenData["access_token"])
	return tokenData["access_token"].(string), nil
}

// Get Gitlab user data from the given authorization code
func GetGitlabUser(code string) (map[string]interface{}, error) {
	var err error

	token, err := GetGitlabAccessToken(code)
	if err != nil {
		log.Printf("Failed to get Gitlab access token, %v", err)
		return nil, err
	}

	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://gitlab.com/api/v4/user", nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	resp, err := client.Do(req)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var userData map[string]interface{}
	err = json.Unmarshal(body, &userData)

	if err != nil {
		log.Printf("Error decoding Github user data, %v", err)
	}

	return userData, nil
}
