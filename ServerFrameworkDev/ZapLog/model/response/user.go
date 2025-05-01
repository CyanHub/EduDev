package response

import "github.com/CyanHub/EduDev/model"

type LoginResponse struct {
	User  *model.User `json:"user"`
	Token string      `json:"token"`
}
