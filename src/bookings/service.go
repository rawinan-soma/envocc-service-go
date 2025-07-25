package bookings

import (
	"envocc-service-go/entities"
	"errors"
	"math/rand"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type BookingService interface {
	GetAllBookings() ([]entities.RoomBooking, error)
	CreateBooking(dto BookingCreate) error
	GetBookingByID(bookingID uint16) (*entities.RoomBooking, error)
	UpdateBooking(bookingID uint16, dto BookingUpdate) error
	DeleteBooking(bookingID uint16) error
}

type bookingService struct {
	repository BookingRepository
}

func NewBookingService(repository BookingRepository) BookingService {
	return &bookingService{
		repository: repository,
	}
}

var (
	ErrRoomReserved     = errors.New("room is not available")
	ErrReservedNotFound = errors.New("not found reservation")
	ErrInvalidDate      = errors.New("invalid start and end date/time")
)

func (s *bookingService) GetAllBookings() ([]entities.RoomBooking, error) {
	var bookings []entities.RoomBooking

	bookings, err := s.repository.FindAll()

	return bookings, err
}

func (s *bookingService) GetBookingByID(bookingID uint16) (*entities.RoomBooking, error) {
	booking, err := s.repository.FindByID(bookingID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrReservedNotFound
	}
	return booking, nil
}

func (s *bookingService) CreateBooking(dto BookingCreate) error {

	if dto.StartDatetime.After(dto.EndDatetime) {
		return ErrInvalidDate
	}

	reservedList, err := s.repository.FindByRoomID(uint8(dto.RoomID))
	if err != nil {
		return err
	}

	for _, req := range reservedList {
		if req.StartDatetime.Before(dto.EndDatetime) && req.EndDatetime.After(dto.StartDatetime) {
			return ErrRoomReserved
		}
	}

	if dto.ConferenceReq != nil {
		src := rand.NewSource(time.Now().UnixNano())
		r := rand.New(src)
		dto.ConferenceReq.Password = strconv.Itoa(r.Intn(900000) + 100000)
	}

	newBooking := ToBookingEntity(dto)
	return s.repository.SaveBooking(newBooking)
}

func (s *bookingService) UpdateBooking(bookingID uint16, dto BookingUpdate) error {

	s.checkBookingExistingByID(bookingID)

	if dto.StartDatetime.After(*dto.EndDatetime) {
		return ErrInvalidDate
	}

	reservedList, err := s.repository.FindByRoomID(uint8(*dto.RoomID))
	if err != nil {
		return err
	}

	for _, req := range reservedList {
		if req.StartDatetime.Before(*dto.EndDatetime) && req.EndDatetime.After(*dto.StartDatetime) {
			return ErrRoomReserved
		}
	}

	updatedBooking := ToBookingEntity(dto)
	return s.repository.UpdateBooking(bookingID, updatedBooking)

}

func (s *bookingService) DeleteBooking(bookingID uint16) error {

	s.checkBookingExistingByID(bookingID)

	if err := s.repository.DeleteBookingByID(bookingID); err != nil {
		return err
	}

	return nil

}

func (s *bookingService) checkBookingExistingByID(bookingID uint16) error {
	_, err := s.repository.FindByID(bookingID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrReservedNotFound
	}

	if err != nil {
		return err
	}
	return nil
}
