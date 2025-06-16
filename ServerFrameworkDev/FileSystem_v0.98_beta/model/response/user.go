package response

import "FileSystem/model"

type LoginResponse struct {
	User  *model.User `json:"user"`
	Token string      `json:"token"`
}
