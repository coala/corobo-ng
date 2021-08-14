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

// Get Github access token from authorization code
func GetGithubAccessToken(code string) (string, error) {
	var err error

	jsonBody, _ := json.Marshal(map[string]string{
		"client_id":     os.Getenv("GITHUB_CLIENT_ID"),
		"client_secret": os.Getenv("GITHUB_CLIENT_SECRET"),
		"code":          code,
		// "redirect_uri":  "http://localhost:3000",
	})

	client := &http.Client{}
	req, _ := http.NewRequest("POST", "https://github.com/login/oauth/access_token", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	log.Printf("Requesting Github for access token, %s", code)

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error getting Github access token, %v", err)
		return "", err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var tokenData map[string]string
	err = json.Unmarshal(body, &tokenData)

	if err != nil {
		log.Printf("Error decoding Github token response, %v", err)
		return "", err
	}

	if _, ok := tokenData["error"]; ok {
		log.Printf("Received an error response from Github while requesting access token")
		return "", error(fmt.Errorf("error response, %v", tokenData))
	}

	log.Printf("Successfully retrieved access token, %s", tokenData["access_token"])
	return tokenData["access_token"], nil
}

// Get Github user data from the given authorization code
func GetGithubUser(code string) (map[string]interface{}, error) {
	var err error

	token, err := GetGithubAccessToken(code)
	if err != nil {
		log.Printf("Failed to get Github access token, %v", err)
		return nil, err
	}

	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://api.github.com/user", nil)
	req.Header.Add("Authorization", fmt.Sprintf("Token %s", token))
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
