package users

import (
	"photo-sharing-api/models"
	"photo-sharing-api/requests"
	"photo-sharing-api/utils"

	"golang.org/x/crypto/bcrypt"
)

func (service *UserService) GetUser(modelUser *models.Users, userId uint) error {
	err := service.DB.First(modelUser, userId).Error
	if err != nil {
		return err
	}
	return nil
}

func (service *UserService) Login(modelUser *models.Users, requestUser *requests.RequestLogin) error {
	service.DB.Where("email = ?", requestUser.Email).First(&modelUser)

	if modelUser.ID == 0 {
		return utils.ErrUserNotFound
	}

	err := bcrypt.CompareHashAndPassword([]byte(modelUser.Password), []byte(requestUser.Password))
	if err != nil {
		return utils.ErrInvalidPassword
	}

	return nil
}
