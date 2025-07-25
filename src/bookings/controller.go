package bookings

import (
	"errors"
	"strconv"

	"github.com/labstack/echo/v4"
)

type bookingController struct {
	service BookingService
}

func NewBookingController(service BookingService) *bookingController {
	return &bookingController{
		service: service,
	}
}

func (c *bookingController) GetAllBookingsHandler(ctx echo.Context) error {
	bookings, err := c.service.GetAllBookings()
	if err != nil {
		return ctx.JSON(500, echo.Map{
			"error": err.Error(),
			"msg":   "something went wrong",
		})
	}

	return ctx.JSON(200, bookings)
}

func (c *bookingController) GetBookingByIDHandler(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(500, echo.Map{"msg": "something went wrong", "err": err.Error()})
	}

	booking, err := c.service.GetBookingByID(uint16(id))
	if errors.Is(err, ErrReservedNotFound) {
		return ctx.JSON(400, echo.Map{"msg": "bad request by user", "err": err.Error()})
	}

	if err != nil {
		return ctx.JSON(500, echo.Map{"msg": "something went wrong", "err": err.Error()})
	}

	return ctx.JSON(200, booking)
}

func (c *bookingController) CreateBookingHandler(ctx echo.Context) error {
	var dto BookingCreate

	if err := ctx.Bind(&dto); err != nil {
		return ctx.JSON(400, echo.Map{"msg": "bad request by user", "error": err.Error()})
	}

	err := c.service.CreateBooking(dto)
	if errors.Is(err, ErrInvalidDate) || errors.Is(err, ErrRoomReserved) {
		return ctx.JSON(400, echo.Map{"msg": "bad request by user", "error": err.Error()})
	}

	if err != nil {
		return ctx.JSON(500, echo.Map{"msg": "something went wrong", "error": err.Error()})
	}

	return ctx.JSON(201, echo.Map{"msg": "booking created", "booking": dto})
}
