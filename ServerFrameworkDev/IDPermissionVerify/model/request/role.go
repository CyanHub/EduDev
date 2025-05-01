package request

type RoleListRequest struct {
	Page     int    `json:"page" form:"page" binding:"required"`
	PageSize int    `json:"pageSize" form:"pageSize" binding:"required"`
	Name     string `json:"name" form:"name" binding:""`
	ParentId int    `json:"parentId" form:"parentId" binding:""`
}

type RoleCreateRequest struct {
	Name     string `json:"name" form:"name" binding:"required"`
	ParentId uint64 `json:"parentId" form:"parentId" binding:""`
}

type RoleUpdateRequest struct {
	Id       uint64 `json:"id" form:"id" binding:"required"`
	Name     string `json:"name" form:"name" binding:"required"`
	ParentId uint64 `json:"parentId" form:"parentId" binding:""`
}
