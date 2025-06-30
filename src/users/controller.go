package users

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type userController struct {
	service UserService
}

func NewUserController(service UserService) *userController {
	return &userController{
		service: service,
	}
}

func (c *userController) GetAllUserHandler(ctx echo.Context) error {
	users, err := c.service.GetAllUsers()

	if err != nil {
		return ctx.JSON(echo.ErrInternalServerError.Code, echo.Map{
			"msg":   echo.ErrInternalServerError.Message,
			"error": err.Error(),
		})
	}

	var response []UserResponse
	for _, user := range users {
		response = append(response, toUserResponse(&user))
	}

	return ctx.JSON(http.StatusOK, response)
}
