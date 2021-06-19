package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/coala/corobo-ng/models"
	"gorm.io/gorm"
)

func GetUser(db *gorm.DB, id string) (*models.User, error) {
	var err error
	user := new(models.User)

	if err = db.Find(&user, id).Error; err != nil {
		log.Println(err)
		return nil, err
	}

	return user, err
}

func GetGithubUser(token string) (map[string]interface{}, error) {
	var err error

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
		log.Printf("Error decoding user data, %v", err)
	}

	return userData, nil
}
