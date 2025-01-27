package users

import (
	"photo-sharing-api/models"
	"photo-sharing-api/requests"
	"photo-sharing-api/utils"

	"golang.org/x/crypto/bcrypt"
)

func (service *UserService) Register(modelUser *models.Users, requestUser *requests.RequestRegister) error {
	service.DB.Where("email = ?", requestUser.Email).First(&modelUser)

	if modelUser.ID != 0 {
		return utils.ErrEmailAlreadyExists
	}

	service.DB.Where("username = ?", requestUser.Username).First(&modelUser)

	if modelUser.ID != 0 {
		return utils.ErrUsernameAlreadyExists
	}

	modelUser.Email = requestUser.Email
	modelUser.Username = requestUser.Username

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(requestUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	modelUser.Password = string(hashedPassword)

	return service.DB.Create(modelUser).Error
}
