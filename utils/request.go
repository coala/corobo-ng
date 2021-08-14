package utils

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func GetRequestBody(c *gin.Context) (map[string]interface{}, error) {
	var err error
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("Unable to read request body, %v", err)
		return nil, err
	}

	var jsonData map[string]interface{}
	err = json.Unmarshal(body, &jsonData)
	if err != nil {
		log.Printf("Error decoding request body, %v", err)
		return nil, err
	}
	return jsonData, nil
}

// Generate a unique token for the user for authorization
func GenerateToken(email string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(email), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Error generating hash, %v", err)
	}

	hasher := md5.New()
	hasher.Write(hash)
	return hex.EncodeToString(hasher.Sum(nil))
}
