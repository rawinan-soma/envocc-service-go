package bookings

import "envocc-service-go/entities"

type BookingService interface {
	GetAllBookings() ([]entities.RoomBooking, error)
}

type bookingService struct {
	repository BookingRepository
}

func NewBookingService(repository BookingRepository) BookingService {
	return &bookingService{
		repository: repository,
	}
}

func (s *bookingService) GetAllBookings() ([]entities.RoomBooking, error) {
	var bookings []entities.RoomBooking

	bookings, err := s.repository.FindAll()

	return bookings, err
}

func (s *bookingService) GetBookingByID(bookingID uint16) (*entities.RoomBooking, error) {
	booking, err := s.repository.FindByID(bookingID)
	return booking, err
}

func (s *bookingService) CreateBooking() {}
