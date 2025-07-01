package users

import (
	"envocc-service-go/entities"

	"github.com/jinzhu/copier"
)

// type UserResponse struct {
// 	ID       int16  `json:"id"`
// 	Username string `json:"username"`
// 	Email    string `json:"email"`
// 	Prefix   string `json:"prefix"`
// 	FnameTH  string `json:"thai_f_name"`
// 	LnameTH  string `json:"thai_l_name"`
// 	Phone    string `json:"phone"`
// 	Line     string `json:"line"`
// 	Role     string `json:"role"`
// 	AvatarID string `json:"avatar_id"`
// }

// func toUserResponse(user *entities.User) UserResponse {
// 	return UserResponse{
// 		ID:       int16(user.ID),
// 		Username: user.Username,
// 		Email:    user.Email,
// 		Prefix:   user.Prefix,
// 		FnameTH:  user.FnameTH,
// 		LnameTH:  user.LnameTH,
// 		Phone:    user.Phone,
// 		Line:     user.Line,
// 		Role:     user.Role,
// 		AvatarID: user.AvatarID,
// 	}
// }

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
