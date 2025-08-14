package bookings

import (
	"envocc-service-go/database"
	"envocc-service-go/src/auth"

	"github.com/labstack/echo/v4"
)

func Wire(group *echo.Group, db database.Database, mw auth.AuthMiddlewareContainer) {
	repository := NewBookingRepository(db)
	service := NewBookingService(repository)
	controller := NewBookingController(service)

	bookingGroup := group.Group("/bookings", mw.JwtAccess)
	bookingGroup.GET("", controller.GetAllBookingsHandler)
	bookingGroup.GET("/:id", controller.GetBookingByIDHandler)
	bookingGroup.POST("", controller.CreateBookingHandler)
	bookingGroup.PATCH("/:id", controller.UpdateBookingHandler)
	bookingGroup.DELETE("/:id", controller.DeleteBookingHandler)

}
