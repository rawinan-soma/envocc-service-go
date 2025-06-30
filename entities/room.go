package entities

type Room struct {
	ID           int16 `gorm:"primaryKey;autoIncrement"`
	Name         string
	ImageURL     string
	Capacity     int16
	HasEquipment bool
}
