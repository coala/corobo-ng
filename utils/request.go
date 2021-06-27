package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
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
