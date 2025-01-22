package responses

import (
	"photo-sharing-api/models"
	"time"

	"github.com/gin-gonic/gin"
)

type ResponseUser struct {
	ID        uint       `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Email     string     `json:"email"`
	Username  string     `json:"username"`
	Password  string     `json:"password"`
	Gendar    string     `json:"gendar"`
	Age       uint8      `json:"age"`
	Birthday  *time.Time `json:"birthday"`
	Address   string     `json:"address"`
	City      string     `json:"city"`
	State     string     `json:"state"`
	Country   string     `json:"country"`
}

func NewResponseUser(context *gin.Context, statusCode int, user models.Users) {
	Response(context, statusCode, ResponseUser{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Username:  user.Username,
		Password:  user.Password,
		Gendar:    user.Gendar,
		Age:       user.Age,
		Birthday:  user.Birthday,
		Address:   user.Address,
		City:      user.City,
		State:     user.State,
		Country:   user.Country,
	})
}

func NewResponseUsers(context *gin.Context, statusCode int, users []models.Users) {
	responseUsers := []ResponseUser{}
	for _, user := range users {
		responseUsers = append(responseUsers, ResponseUser{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Username:  user.Username,
			Password:  user.Password,
			Gendar:    user.Gendar,
			Age:       user.Age,
			Birthday:  user.Birthday,
			Address:   user.Address,
			City:      user.City,
			State:     user.State,
			Country:   user.Country,
		})
	}

	Response(context, statusCode, responseUsers)
}
