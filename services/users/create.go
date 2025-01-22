package users

import (
	"photo-sharing-api/models"
	"photo-sharing-api/requests"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func (service *UserService) Create(modelUser *models.Users, requestUser *requests.RequestUser) error {
	modelUser.FirstName = requestUser.FirstName
	modelUser.LastName = requestUser.LastName
	modelUser.Email = requestUser.Email
	modelUser.Username = requestUser.Username

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(requestUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	modelUser.Password = string(hashedPassword)

	modelUser.Gender = requestUser.Gender
	modelUser.Age = requestUser.Age

	birthday, err := time.Parse(time.RFC3339, requestUser.Birthday)
	if err != nil {
		return err
	}

	modelUser.Birthday = &birthday
	modelUser.Address = requestUser.Address
	modelUser.City = requestUser.City
	modelUser.State = requestUser.State
	modelUser.Country = requestUser.Country

	return service.DB.Create(modelUser).Error
}
