package response

import "ServerFramework/model"

type LoginResponse struct {
	User  *model.User `json:"user"`
	Token string      `json:"token"`
}
