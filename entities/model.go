package entities

func GetAllModels() []any {
	return []any{
		&User{}, &Room{}, &Position{}, &PositionLevel{}, &Group{},
	}
}
