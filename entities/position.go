package entities

type Position struct {
	ID    int16 `gorm:"primaryKey;autoIncrement"`
	Name  string
	Users []User
}
