package services

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/coala/corobo-ng/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func generateToken(email string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(email), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Error generating hash, %v", err)
	}

	hasher := md5.New()
	hasher.Write(hash)
	return hex.EncodeToString(hasher.Sum(nil))
}

func GetUser(db *gorm.DB, id string) (*models.User, error) {
	var err error
	user := new(models.User)

	if err = db.Find(&user, id).Error; err != nil {
		log.Println(err)
		return nil, err
	}

	return user, err
}

func GetUserByEmail(db *gorm.DB, email string) (*models.User, error) {
	var err error
	user := new(models.User)

	if err = db.Find(&user, "email = ?", email).Error; err != nil {
		log.Println(err)
		return nil, err
	}

	return user, err
}

func GetUserByProviderId(db *gorm.DB, providerId int64) (*models.User, error) {
	var err error
	user := new(models.User)

	if err = db.Find(&user, "providerId = ?", providerId).Error; err != nil {
		log.Println(err)
		return nil, err
	}

	return user, err
}

func CreateUser(db *gorm.DB, userData map[string]interface{}) (*models.User, error) {
	var err error
	user := &models.User{
		Name:       userData["name"].(string),
		Email:      userData["email"].(string),
		ProviderId: userData["id"].(int64),
		Token:      generateToken(userData["email"].(string)),
	}

	if err = db.Create(&user).Error; err != nil {
		log.Printf("Error creating new user, %v", err)
		return nil, err
	}

	return user, nil
}

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
		log.Printf("Error decoding Github token response, %v", string(bytes.Replace(body, []byte("\r"), []byte("\r\n"), -1)))
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		log.Printf("Received non-200 response from Github while requesting access token")
		return "", error(fmt.Errorf("non-200 response %v", body))
	}

	log.Printf("Successfully retrieved access token, %s", tokenData["access_token"])
	return tokenData["access_token"], nil
}

func GetGithubUser(code string) (map[string]interface{}, error) {
	var err error

	token, err := GetGithubAccessToken(code)
	if err != nil {
		log.Printf("Failed to get Github access token %v", err)
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
