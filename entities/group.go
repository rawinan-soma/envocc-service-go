package entities

type Group struct {
	ID    int16 `gorm:"primaryKey;autoIncrement"`
	Name  string
	Users []User
}
