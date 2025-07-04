package users

import (
	"envocc-service-go/entities"

	"github.com/jinzhu/copier"
)

type UserCreate struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	Email           string `json:"email"`
	Prefix          string `json:"prefix"`
	FnameTH         string `json:"thai_f_name"`
	LnameTH         string `json:"thai_l_name"`
	FnameEn         string `json:"eng_f_name"`
	LnameEn         string `json:"eng_l_name"`
	Phone           string `json:"phone"`
	Line            string `json:"line_id"`
	AvatarID        string `json:"avatar_id"`
	GroupID         int16  `json:"group_id"`
	PositionID      int16  `json:"position_id"`
	PositionLevelID int16  `json:"position_level_id"`
}

func ToUserEntity(dto UserCreate) *entities.User {
	// return &entities.User{
	// 	Username: dto.Username,
	// 	Password: dto.Password,
	// 	Email:    dto.Email,
	// 	FnameTH:  dto.FnameTH,
	// 	LnameTH:  dto.LnameTH,
	// 	FnameEn:  dto.FnameEn,
	// 	LnameEn:  dto.LnameEn,
	// 	Phone:    dto.Phone,
	// 	Line:     dto.Line,
	// }
	user := &entities.User{}
	copier.Copy(user, &dto)
	return user
}
