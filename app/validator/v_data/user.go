package v_data

type VRoleList struct {
	RoleName string `json:"role_name" form:"role_name"`
	State    uint8  `json:"state" form:"state" binding:"numeric"`
}
type VRoleCreate struct {
	RoleName    string   `json:"role_name" form:"role_name" binding:"required"`
	Permissions []string `json:"permissions" binding:"required"`
}
type VRoleUpdate struct {
	Id          int64    `json:"id" form:"id" binding:"required"`
	RoleName    string   `json:"role_name" form:"role_name" binding:"required"`
	State       uint8    `json:"state" form:"state" binding:"numeric"`
	Permissions []string `json:"permissions" binding:"required"`
}

type VUserList struct {
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	State    uint8  `json:"state" form:"state" binding:"numeric"`
	Pagination
}

type VUserCreate struct {
	Username string `json:"username" form:"username" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Mobile   string `json:"mobile" form:"mobile"`
	State    uint8  `json:"state" form:"state" binding:"numeric"`
	RoleId   int64  `json:"role_id" form:"role_id" binding:"required,numeric"`
	Pass     string `json:"pass" form:"pass"`
}

type VUserUpdate struct {
	Id int64 `json:"id" form:"id" binding:"required,numeric"`
	VUserCreate
}

type VLogin struct {
	Email string `json:"email" binding:"required,email"`
	Pass  string `json:"pass" binding:"required,pass"`
}

type VPermissionCreate struct {
	Permission string `json:"permission" binding:"required"`
	PName      string `json:"p_name" binding:"required"`
	Method     string `json:"method" binding:"required"`
	Pid        int64  `json:"pid" binding:"numeric"`
}

type VPermissionUpdate struct {
	Id int64 `json:"id" binding:"required,numeric"`
	VPermissionCreate
}

type VSsoLoginData struct {
	Ticket string `json:"ticket" binding:"required,alphanum"`
}

type VRolePermissions struct {
	Id int64 `form:"id" binding:"required,numeric"`
}
