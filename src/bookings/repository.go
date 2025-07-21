package bookings

import (
	"envocc-service-go/database"
	"envocc-service-go/entities"

	"gorm.io/gorm"
)

type BookingRepository interface {
	FindAll() ([]entities.RoomBooking, error)
	FindByID(bookingID uint16) (*entities.RoomBooking, error)
	DeleteBookingByID(bookingID uint16) error
	SaveBooking(booking *entities.RoomBooking) error
}

type bookingRepository struct {
	db *gorm.DB
}

func NewBookingRepository(db database.Database) BookingRepository {
	return &bookingRepository{
		db: db.GetDb(),
	}
}

func (r *bookingRepository) FindAll() ([]entities.RoomBooking, error) {
	var bookings []entities.RoomBooking
	return bookings, r.db.Preload("ConferenceReq").Find(&bookings).Error
}

func (r *bookingRepository) FindByID(bookingID uint16) (*entities.RoomBooking, error) {
	var booking entities.RoomBooking
	err := r.db.Where("id = ?", bookingID).Find(&booking).Error

	return &booking, err
}

func (r *bookingRepository) SaveBooking(booking *entities.RoomBooking) error {
	return r.db.Create(booking).Error
}

func (r *bookingRepository) DeleteBookingByID(bookingID uint16) error {
	return r.db.Delete(&entities.RoomBooking{}, 10).Error
}

func (r *bookingRepository) UpdateBookig(bookingID uint16, newBooking *entities.RoomBooking) error {
	return r.db.Model(&entities.RoomBooking{}).Where("id = ?", bookingID).Updates(newBooking).Error
}
