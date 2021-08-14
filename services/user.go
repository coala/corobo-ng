package services

import (
	"log"

	"github.com/coala/corobo-ng/models"
	"github.com/coala/corobo-ng/utils"
	"gorm.io/gorm"
)

// Get user object for given id
func GetUser(db *gorm.DB, id string) (*models.User, error) {
	var err error
	user := new(models.User)

	if err = db.Find(&user, id).Error; err != nil {
		log.Println(err)
		return nil, err
	}

	return user, err
}

// Get user object by email
func GetUserByEmail(db *gorm.DB, email string) (*models.User, error) {
	var err error
	user := new(models.User)

	if err = db.First(&user, "email = ?", email).Error; err != nil {
		log.Println(err)
		return nil, err
	}

	return user, err
}

// Get user object by given provider id
func GetUserByProviderId(db *gorm.DB, providerId int64) (*models.User, error) {
	var err error
	user := new(models.User)

	if err = db.Find(&user, "providerId = ?", providerId).Error; err != nil {
		log.Println(err)
		return nil, err
	}

	return user, err
}

// Get user object by given token
func GetUserByToken(db *gorm.DB, token string) (*models.User, error) {
	var err error
	user := new(models.User)

	if err = db.Find(&user, "token = ?", token).Error; err != nil {
		log.Println(err)
		return nil, err
	}

	return user, err
}

// Create new user
func CreateUser(db *gorm.DB, userData map[string]interface{}) (*models.User, error) {
	var err error
	user := &models.User{
		Name:       userData["name"].(string),
		Email:      userData["email"].(string),
		ProviderId: int64(userData["id"].(float64)),
		Token:      utils.GenerateToken(userData["email"].(string)),
	}

	if err = db.Create(&user).Error; err != nil {
		log.Printf("Error creating new user, %v", err)
		return nil, err
	}

	return user, nil
}
