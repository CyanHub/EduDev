package request

type RoleListRequest struct {
	Page     int    `json:"page" form:"page" binding:"required,gt=0" chinese:"页码"`
	PageSize int    `json:"pageSize" form:"pageSize" binding:"required,gt=0" chinese:"每页数量"`
	Name     string `json:"name" form:"name" binding:"" chinese:"角色名称"`
	ParentId uint64 `json:"parentId" form:"parentId" binding:"gte=0" chinese:"父级角色Id"`
}

type RoleCreateRequest struct {
	Name     string `json:"name" form:"name" binding:"required" chinese:"角色名称"`
	ParentId uint64 `json:"parentId" form:"parentId" binding:"gte=0" chinese:"父级角色Id"`
}

type RoleUpdateRequest struct {
	Id       uint64 `json:"id" form:"id" binding:"required" chinese:"角色Id"`
	Name     string `json:"name" form:"name" binding:"required" chinese:"角色名称"`
	ParentId uint64 `json:"parentId" form:"parentId" binding:"gte=0" chinese:"父级角色Id"`
}

type AssignPermissions struct {
	RoleID      uint64   `json:"roleId" form:"roleId" binding:"required" chinese:"角色ID"`
	Permissions []string `json:"permissions" form:"permissions" binding:"required" chinese:"权限列表"`
}

type CheckPermission struct {
	UserID     uint64 `json:"userId" form:"userId" binding:"required" chinese:"用户ID"`
	Permission string `json:"permission" form:"permission" binding:"required" chinese:"权限"`
}
