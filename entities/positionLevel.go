package entities

type PositionLevel struct {
	ID    int16 `gorm:"primaryKey;autoIncrement"`
	Name  string
	Users []User `gorm:"foreignKey:PositionLevelID"`
}

func (PositionLevel) TableName() string {
	return "position_level"
}
