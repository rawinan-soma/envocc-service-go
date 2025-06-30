package users

import "envocc-service-go/entities"

type UserResponse struct {
	ID       int16  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Prefix   string `json:"prefix"`
	FnameTH  string `json:"thai_f_name"`
	LnameTH  string `json:"thai_l_name"`
	Phone    string `json:"phone"`
	Line     string `json:"line"`
	Role     string `json:"role"`
	AvatarID string `json:"avatar_id"`
}

func toUserResponse(user *entities.User) UserResponse {
	return UserResponse{
		ID:       int16(user.ID),
		Username: user.Username,
		Email:    user.Email,
		Prefix:   user.Prefix,
		FnameTH:  user.FnameTH,
		LnameTH:  user.LnameTH,
		Phone:    user.Phone,
		Line:     user.Line,
		Role:     user.Role,
		AvatarID: user.AvatarID,
	}
}
