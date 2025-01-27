package users

import (
	"photo-sharing-api/models"
	"photo-sharing-api/requests"
	"time"
)

func (service *UserService) Update(userId uint, modelUser *models.Users, requestUser *requests.RequestUpdateUser) error {
	err := service.DB.First(modelUser, userId).Error
	if err != nil {
		return err
	}

	modelUser.FirstName = requestUser.FirstName
	modelUser.LastName = requestUser.LastName
	modelUser.Username = requestUser.Username

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
	modelUser.Avatar = requestUser.Avatar

	return service.DB.Save(modelUser).Error
}
