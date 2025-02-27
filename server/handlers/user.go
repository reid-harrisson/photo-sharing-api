package handlers

import (
	"net/http"
	"photo-sharing-api/models"
	"photo-sharing-api/requests"
	"photo-sharing-api/responses"
	"photo-sharing-api/server"
	"photo-sharing-api/services/users"
	"photo-sharing-api/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Server  *server.Server
	Service *users.UserService
}

func NewUserHandler(server *server.Server) *UserHandler {
	return &UserHandler{
		Server:  server,
		Service: users.NewUserService(server.DB),
	}
}

// Users godoc
// @Summary Register user
// @Schemes
// @Description Register user
// @Tags Users
// @Accept json
// @Produce json
// @Param request body requests.RequestRegister true "User registration data"
// @Success 200 {object} responses.ResponseUser
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /users/register [post]
func (handler *UserHandler) Register(context *gin.Context) {
	request := requests.RequestRegister{}

	// Bind the incoming JSON to the RequestRegister struct
	if err := context.ShouldBindJSON(&request); err != nil {
		responses.ErrorResponse(context, http.StatusBadRequest, utils.MsgInvalidRequestData)
		return
	}

	user := models.Users{}

	// Create the user using the UserService
	if err := handler.Service.Register(&user, &request); err != nil {
		if err == utils.ErrEmailAlreadyExists {
			responses.ErrorResponse(context, http.StatusBadRequest, utils.MsgEmailAlreadyExists)
		} else if err == utils.ErrUsernameAlreadyExists {
			responses.ErrorResponse(context, http.StatusBadRequest, utils.MsgUsernameAlreadyExists)
		} else {
			responses.ErrorResponse(context, http.StatusInternalServerError, utils.MsgFailedToUpdateUser)
		}
		return
	}

	responses.NewResponseUser(context, http.StatusOK, user)
}

// Users godoc
// @Summary Login user
// @Schemes
// @Description Login user
// @Tags Users
// @Accept json
// @Produce json
// @Param request body requests.RequestLogin true "User login data"
// @Success 200 {object} responses.ResponseUser
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /users/login [post]
func (handler *UserHandler) Login(context *gin.Context) {
	request := requests.RequestLogin{}

	if err := context.ShouldBindJSON(&request); err != nil {
		responses.ErrorResponse(context, http.StatusBadRequest, utils.MsgInvalidRequestData)
		return
	}

	user := models.Users{}

	if err := handler.Service.Login(&user, &request); err != nil {
		if err == utils.ErrUserNotFound {
			responses.ErrorResponse(context, http.StatusBadRequest, utils.MsgUserNotFound)
		} else if err == utils.ErrInvalidPassword {
			responses.ErrorResponse(context, http.StatusBadRequest, utils.MsgInvalidPassword)
		} else {
			responses.ErrorResponse(context, http.StatusInternalServerError, utils.MsgFailedToLogin)
		}
		return
	}

	responses.NewResponseUser(context, http.StatusOK, user)
}

// Users godoc
// @Summary Login user by username
// @Schemes
// @Description Login user by username
// @Tags Users
// @Accept json
// @Produce json
// @Param request body requests.RequestLoginByUsername true "User login data by username"
// @Success 200 {object} responses.ResponseUser
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /users/login-by-username [post]
func (handler *UserHandler) LoginByUsername(context *gin.Context) {
	request := requests.RequestLoginByUsername{}

	if err := context.ShouldBindJSON(&request); err != nil {
		responses.ErrorResponse(context, http.StatusBadRequest, utils.MsgInvalidRequestData)
		return
	}

	user := models.Users{}

	if err := handler.Service.LoginByUsername(&user, &request); err != nil {
		if err == utils.ErrUserNotFound {
			responses.ErrorResponse(context, http.StatusBadRequest, utils.MsgUserNotFound)
		} else if err == utils.ErrInvalidPassword {
			responses.ErrorResponse(context, http.StatusBadRequest, utils.MsgInvalidPassword)
		} else {
			responses.ErrorResponse(context, http.StatusInternalServerError, utils.MsgFailedToLogin)
		}
		return
	}

	responses.NewResponseUser(context, http.StatusOK, user)
}

// Users godoc
// @Summary Update user
// @Schemes
// @Description Update user
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param request body requests.RequestUpdateUser true "User update data"
// @Success 200 {object} responses.ResponseUser
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /users/update/{id} [put]
func (handler *UserHandler) Update(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 10, 64)
	if err != nil {
		responses.ErrorResponse(context, http.StatusBadRequest, utils.MsgInvalidID)
		return
	}

	request := requests.RequestUpdateUser{}

	if err := context.ShouldBindJSON(&request); err != nil {
		responses.ErrorResponse(context, http.StatusBadRequest, utils.MsgInvalidRequestData)
		return
	}

	user := models.Users{}

	if err := handler.Service.Update(uint(id), &user, &request); err != nil {
		if err == utils.ErrUserNotFound {
			responses.ErrorResponse(context, http.StatusBadRequest, utils.MsgUserNotFound)
		} else if err == utils.ErrInvalidPassword {
			responses.ErrorResponse(context, http.StatusBadRequest, utils.MsgInvalidPassword)
		} else {
			responses.ErrorResponse(context, http.StatusInternalServerError, utils.MsgFailedToUpdateUser)
		}
		return
	}

	responses.NewResponseUser(context, http.StatusOK, user)
}
