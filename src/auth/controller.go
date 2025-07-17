package auth

import (
	"envocc-service-go/entities"
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
	var dto UserCreate

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

func (c *authenController) LoginHandler(ctx echo.Context) error {
	userCtx := ctx.Get("user")
	user, ok := userCtx.(*entities.User)
	if !ok {
		return ctx.JSON(500, echo.Map{
			"msg": "cannot get user",
		})
	}

	accessTokenCookie, err := c.service.GetCookieWithAccessToken(user.ID)
	refreshTokenCookie, refreshToken, err := c.service.GetCookieWithRefreshToken(user.ID)

	if err != nil {
		return ctx.JSON(500, echo.Map{
			"msg":   "cannot generate cookies",
			"error": err.Error(),
		})
	}

	if err := c.service.SetCurrentRefreshToken(refreshToken, user.ID); err != nil {
		return ctx.JSON(500, echo.Map{"msg": "something went wrong", "error": err.Error()})
	}

	ctx.SetCookie(accessTokenCookie)
	ctx.SetCookie(refreshTokenCookie)

	return ctx.JSON(200, echo.Map{"msg": "login successful"})
}

func (c *authenController) RefreshTokenHandler(ctx echo.Context) error {
	userCtx := ctx.Get("user")
	user, ok := userCtx.(*entities.User)
	if !ok {
		return ctx.JSON(500, echo.Map{"msg": "something went wrong"})
	}

	accessTokenCookie, err := c.service.GetCookieWithAccessToken(user.ID)
	if err != nil {
		return ctx.JSON(500, echo.Map{"msg": "cannot generate new cookie", "error": err.Error()})
	}

	ctx.SetCookie(accessTokenCookie)
	return ctx.JSON(200, echo.Map{"msg": "token refreshed"})
}

func (c *authenController) LogoutHandler(ctx echo.Context) error {
	userCtx := ctx.Get("user")
	user, ok := userCtx.(*entities.User)
	if !ok {
		return ctx.JSON(500, echo.Map{"msg": "cannot get user context"})
	}

	if err := c.service.RemoveRefreshToken(user.ID); err != nil {
		return ctx.JSON(500, echo.Map{"msg": "something went wrong", "error": err.Error()})
	}

	cookies := c.service.GetLogoutCookie()
	for _, cookie := range cookies {
		ctx.SetCookie(cookie)
	}

	return ctx.JSON(200, echo.Map{"msg": "logout succesful"})
}
