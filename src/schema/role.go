package schema

type RoleForm struct{}

type RoleSchema struct {
	Name string `json:"name" binding:"required"`
}
