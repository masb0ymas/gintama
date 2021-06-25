package schema

type RoleByIdSchema struct {
	ID int `uri:"id" binding:"required"`
}

type RoleSchema struct {
	Name string `json:"name" binding:"required"`
}
