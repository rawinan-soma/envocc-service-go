package auth

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// Helper function to parse JWT token

func LocalAuthMiddleware(service AuthenService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var credential struct {
				Username string `json:"username" validate:"required"`
				Password string `json:"password" validate:"required"`
			}

			fmt.Println(c.Request().Header.Get("Content-Type"))

			if err := c.Bind(&credential); err != nil {
				return echo.ErrBadRequest
			}
			validator := validator.New()
			if err := validator.Struct(credential); err != nil {
				return echo.ErrBadRequest
			}
			user, err := service.GetAuthenticatedUser(credential.Username, credential.Password)
			if err != nil {
				return c.JSON(echo.ErrUnauthorized.Code, echo.Map{"msg": "invalid credential", "from": "controller"})
			}
			c.Set("user", user)
			return next(c)
		}
	}
}

func JwtAccessMiddleware(service AuthenService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cookie, err := c.Cookie("Authentication")
			if err != nil {
				return c.JSON(echo.ErrUnauthorized.Code, echo.Map{"msg": "unauthorized"})
			}

			claims, err := service.ParseAccessToken(cookie.Value)
			if err != nil {
				return c.JSON(echo.ErrUnauthorized.Code, echo.Map{"msg": "invalid token"})
			}

			user, err := service.GetUserById(uint16(claims.UserID))
			if err != nil {
				return c.JSON(echo.ErrUnauthorized.Code, echo.Map{"msg": "invalid token"})
			}

			user.Password = ""
			user.HashedRefreshToken = nil

			c.Set("user", &user)
			return next(c)
		}
	}
}

func JwtRefreshMiddleware(service AuthenService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cookie, err := c.Cookie("Refresh")
			if err != nil {
				return c.JSON(401, echo.Map{"msg": "invalid token"})
			}

			claims, err := service.ParseRefreshToken(cookie.Value)
			if err != nil {
				return c.JSON(401, echo.Map{"msg": "invalid token"})
			}

			user, err := service.GetUserFromRefreshToken(cookie.Value, uint16(claims.UserID))
			if err != nil {
				return c.JSON(401, echo.Map{"msg": "invalid token"})
			}
			c.Set("user", &user)
			return next(c)
		}
	}
}

type AuthMiddlewareContainer struct {
	Local      echo.MiddlewareFunc
	JwtAccess  echo.MiddlewareFunc
	JwtRefresh echo.MiddlewareFunc
}

func NewMiddlewareConatiner(service AuthenService) *AuthMiddlewareContainer {
	return &AuthMiddlewareContainer{
		Local:      LocalAuthMiddleware(service),
		JwtAccess:  JwtAccessMiddleware(service),
		JwtRefresh: JwtRefreshMiddleware(service),
	}
}
