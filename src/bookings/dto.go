package bookings

import (
	"envocc-service-go/entities"
	"time"

	"github.com/jinzhu/copier"
)

type BookingCreate struct {
	UserID        int16                `json:"user_id"`
	RoomID        int16                `json:"room_id"`
	Title         string               `json:"meeting_title"`
	Attendees     int16                `json:"attendees"`
	StartDatetime time.Time            `json:"start_datetime"`
	EndDatetime   time.Time            `json:"end_datetime"`
	NeedEquipment bool                 `json:"need_equipment"`
	Notes         *string              `json:"notes"`
	ConferenceReq *ConferenceReqCreate `json:"conference_req,omitempty"`
}

type ConferenceReqCreate struct {
	Title     string `json:"meeting_title"`
	Password  string `json:"password,omitempty"`
	Url       string `json:"meeting_url"`
	Equipment string `json:"equipment"`
	Host      bool   `json:"host"`
}

func ToBookingEntity[T BookingCreate | BookingUpdate](dto T) *entities.RoomBooking {
	booking := &entities.RoomBooking{}
	copier.CopyWithOption(booking, &dto, copier.Option{IgnoreEmpty: true})

	switch v := any(dto).(type) {
	case BookingCreate:
		if v.ConferenceReq != nil {
			booking.ConferenceReq = &entities.ConferenceRequest{}
			copier.CopyWithOption(booking.ConferenceReq, v.ConferenceReq, copier.Option{IgnoreEmpty: true})
		}
	case BookingUpdate:
		if v.ConferenceReq != nil {
			booking.ConferenceReq = &entities.ConferenceRequest{}
			copier.CopyWithOption(booking.ConferenceReq, v.ConferenceReq, copier.Option{IgnoreEmpty: true})
		}
	}

	// if dto.ConferenceReq != nil {
	// 	booking.ConferenceReq = &entities.ConferenceRequest{}
	// 	copier.CopyWithOption(booking.ConferenceReq, &dto.ConferenceReq, copier.Option{IgnoreEmpty: true})
	// }

	return booking
}

type BookingUpdate struct {
	UserID        *int16               `json:"user_id,omitempty"`
	RoomID        *int16               `json:"room_id,omitempty"`
	Title         *string              `json:"meeting_title,omitempty"`
	Attendees     *int16               `json:"attendees,omitempty"`
	StartDatetime *time.Time           `json:"start_datetime,omitempty"`
	EndDatetime   *time.Time           `json:"end_datetime,omitempty"`
	NeedEquipment *bool                `json:"need_equipment,omitempty"`
	Notes         *string              `json:"notes,omitempty"`
	ConferenceReq *ConferenceReqUpdate `json:"conference_req,omitempty"`
}

type ConferenceReqUpdate struct {
	Title     *string `json:"meeting_title,omitempty"`
	Password  *string `json:"password,omitempty"`
	Url       *string `json:"meeting_url,omitempty"`
	Equipment *string `json:"equipment,omitempty"`
	Host      bool    `json:"host"`
}
