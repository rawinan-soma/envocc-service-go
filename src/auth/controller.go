package auth

import (
	"envocc-service-go/src/users"
	"errors"

	"github.com/labstack/echo/v4"
)

type authenController struct {
	service AuthenService
}

func NewAuthenController(service AuthenService) *authenController {
	return &authenController{
		service: service,
	}
}

func (c *authenController) CreateUserHandler(ctx echo.Context) error {
	var dto users.UserCreate

	if err := ctx.Bind(&dto); err != nil {
		return ctx.JSON(400, echo.Map{
			"msg":   "bad request by user",
			"error": err.Error(),
		})
	}

	err := c.service.CreateUser(dto)
	if errors.Is(err, ErrExistingUser) {
		return ctx.JSON(400, echo.Map{"msg": "bad request by user", "error": err.Error()})
	}

	if err != nil {
		return ctx.JSON(500, echo.Map{"msg": "something went wrong", "error": err.Error()})
	}

	return ctx.JSON(201, echo.Map{"msg": "user created successful"})
}
