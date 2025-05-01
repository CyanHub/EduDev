package request

type UserLoginRequest struct {
	Username string `json:"username" binding:"required" chinese:"用户名"`
	Password string `json:"password" binding:"required" chinese:"密码"`
}

type UserRegisterRequest struct {
	ID       uint64 `json:"id" binding:"gte=0" chinese:"用户Id"`
	Username string `json:"username" binding:"required,min=3,max=20,usernameUnique" chinese:"用户名"`
	Password string `json:"password" binding:"required,min=6,max=20" chinese:"密码"`
	NickName string `json:"nickname" binding:"" chinese:"昵称"`
}

type UserListRequest struct {
	Page     int    `json:"page" binding:"required,gt=0" chinese:"页码"`
	PageSize int    `json:"pageSize" binding:"required,gt=0" chinese:"每页数量"`
	Username string `json:"username" binding:"min=3,max=20" chinese:"用户名"`
	NickName string `json:"nickname" binding:"" chinese:"昵称"`
	Status   int    `json:"status" binding:"gt=0,lte=2" chinese:"状态"`
	RoleId   uint64 `json:"roleId" binding:"gte=0" chinese:"角色Id"`
}
