package entities

import "time"

type User struct {
	ID              uint16          `gorm:"primaryKey;autoIncrement" json:"id"`
	Username        string          `gorm:"unique" json:"username"`
	Password        string          `json:"-"`
	Email           string          `gorm:"unique" json:"email"`
	Prefix          string          `json:"prefix"`
	FnameTH         string          `gorm:"column:thai_f_name" json:"thai_f_name"`
	LnameTH         string          `gorm:"column:thai_l_name" json:"thai_l_name"`
	FnameEn         string          `gorm:"column:eng_f_name" json:"eng_f_name"`
	LnameEn         string          `gorm:"column:eng_l_name" json:"eng_l_name"`
	Phone           string          `json:"phone"`
	Line            string          `gorm:"column:line_id" json:"line_id"`
	Role            *string         `gorm:"default:user" json:"role"`
	Status          *bool           `gorm:"default:false" json:"status"`
	AvatarID        string          `json:"avatar_id"`
	GroupID         int16           `json:"group_id"`
	PositionID      int16           `json:"position_id"`
	PositionLevelID int16           `json:"position_level_id"`
	RefreshToken    *string         `gorm:"column:refreshToken" json:"-"`
	CreatedAt       time.Time       `gorm:"default:CURRENT_TIMESTAMP" json:"-"`
	Group           Group           `json:"group"`
	Position        Position        `json:"position"`
	PositionLevel   PositionLevel   `json:"position_level"`
	RoomBooking     []RoomBooking   `json:"room_booking"`
	ConferenceReq   []ConferenceReq `json:"conference_req"`
}

type Group struct {
	ID    int16  `gorm:"primaryKey;autoIncrement" json:"id"`
	Name  string `json:"name"`
	Users []User `json:"-"`
}

type Position struct {
	ID    int16  `gorm:"primaryKey;autoIncrement" json:"id"`
	Name  string `json:"name"`
	Users []User `json:"-"`
}

type PositionLevel struct {
	ID    int16  `gorm:"primaryKey;autoIncrement" json:"id"`
	Name  string `json:"name"`
	Users []User `json:"-"`
}

type Room struct {
	ID           int16         `gorm:"primaryKey;autoIncrement" json:"id"`
	Name         string        `json:"name"`
	ImageURL     string        `json:"image_url"`
	Capacity     int16         `json:"capacity"`
	HasEquipment bool          `json:"has_equipment"`
	RoomBooking  []RoomBooking `json:"room_booking"`
}

type RoomBooking struct {
	ID            int16         `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UserID        int16         `json:"user_id"`
	RoomID        int16         `json:"room_id"`
	Title         string        `gorm:"column:meeting_title" json:"meeting_title" `
	Attendees     int16         `json:"attendees"`
	StartDatetime time.Time     `gorm:"column:start_datetime" json:"start_datetime"`
	EndDatetime   time.Time     `gorm:"column:end_datetime" json:"end_datetime"`
	NeedEquipment bool          `json:"need_equipment"`
	Notes         string        `json:"notes"`
	ConferenceReq ConferenceReq `json:"conference_req"`
}

type ConferenceReq struct {
	ID            int16     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	RoomBookingID int16     `json:"room_booking_id"`
	UserID        int16     `json:"user_id"`
	Title         string    `gorm:"column:meeting_title" json:"meeting_title"`
	Password      string    `gorm:"column:meeting_password" json:"meeting_password"`
	StartDatetime time.Time `json:"start_datetime"`
	EndDatetime   time.Time `json:"end_datetime"`
	Equipment     string    `json:"equipment"`
	ConfApp       *string   `json:"conf_app"`
	ConfUsername  *string   `json:"conf_username"`
}

type ConferenceRes struct {
	ID              int16  `gorm:"column:id;type:int;primaryKey;autoIncrement" json:"id"`
	ConferenceReqID int16  `json:"conference_req_id"`
	Title           string `gorm:"column:meeting_title" json:"meeting_title"`
	Password        string `gorm:"column:meeting_password" json:"meeting_password"`
	Url             string `gorm:"column:meeting_url" json:"meeting_url"`
	Number          string `gorm:"column:meeting_number" json:"meeting_number"`
}

func (PositionLevel) TableName() string {
	return "position_level"
}

// FUNCTION to get all models

func GetAllModels() []any {
	return []any{
		&User{}, &Room{}, &Position{}, &PositionLevel{}, &Group{},
	}
}
