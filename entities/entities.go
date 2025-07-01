package entities

import "time"

type User struct {
	ID                 uint16        `gorm:"primaryKey;autoIncrement" json:"id"`
	Username           string        `gorm:"unique" json:"username"`
	Password           string        `json:"-"`
	Email              string        `gorm:"unique" json:"email"`
	Prefix             string        `json:"prefix"`
	FnameTH            string        `gorm:"column:thai_f_name" json:"thai_f_name"`
	LnameTH            string        `gorm:"column:thai_l_name" json:"thai_l_name"`
	FnameEn            string        `gorm:"column:eng_f_name" json:"eng_f_name"`
	LnameEn            string        `gorm:"column:eng_l_name" json:"eng_l_name"`
	Phone              string        `json:"phone"`
	Line               string        `gorm:"column:line_id" json:"line_id"`
	Role               *string       `gorm:"default:user" json:"role"`
	Status             *bool         `gorm:"default:false" json:"status"`
	AvatarID           string        `json:"avatar_id"`
	GroupID            int16         `json:"group_id"`
	PositionID         int16         `json:"position_id"`
	PositionLevelID    int16         `json:"position_level_id"`
	HashedRefreshToken *string       `gorm:"column:hashedRefreshToken" json:"-"`
	CreatedAt          time.Time     `gorm:"default:CURRENT_TIMESTAMP" json:"-"`
	Group              Group         `json:"group"`
	Position           Position      `json:"position"`
	PositionLevel      PositionLevel `json:"position_level"`
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
	ID           int16 `gorm:"primaryKey;autoIncrement"`
	Name         string
	ImageURL     string
	Capacity     int16
	HasEquipment bool
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
