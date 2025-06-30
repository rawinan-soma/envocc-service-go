package entities

import "time"

type User struct {
	ID                 uint16 `gorm:"primaryKey;autoIncrement"`
	Username           string `gorm:"unique"`
	Password           string
	Email              string `gorm:"unique"`
	Prefix             string
	FnameTH            string `gorm:"column:thai_f_name"`
	LnameTH            string `gorm:"column:thai_l_name"`
	FnameEn            string `gorm:"column:eng_f_name"`
	LnameEn            string `gorm:"column:eng_l_name"`
	Phone              string
	Line               string `gorm:"column:line_id"`
	Role               string `gorm:"default:user"`
	Status             bool   `gorm:"default:false"`
	AvatarID           string
	GroupID            int16
	PositionID         int16
	PositionLevelID    int16
	HashedRefreshToken *string   `gorm:"column:hashedRefreshToken"`
	CreatedAt          time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

type Group struct {
	ID    int16 `gorm:"primaryKey;autoIncrement"`
	Name  string
	Users []User
}

type Position struct {
	ID    int16 `gorm:"primaryKey;autoIncrement"`
	Name  string
	Users []User
}

type PositionLevel struct {
	ID    int16 `gorm:"primaryKey;autoIncrement"`
	Name  string
	Users []User `gorm:"foreignKey:PositionLevelID"`
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
