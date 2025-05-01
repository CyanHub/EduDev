package request

type UserLoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserRegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	NickName string `json:"nickname" binding:""`
}

type UserListRequest struct {
	Page     int    `json:"page" binding:"required"`
	PageSize int    `json:"pageSize" binding:"required"`
	Username string `json:"username" binding:""`
	NickName string `json:"nickname" binding:""`
	Status   int    `json:"status" binding:""`
	RoleId   uint64 `json:"roleId" binding:""`
}
