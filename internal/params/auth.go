package params

type RegisterParam struct {
	Name     string `json:"name" binding:"required,max=10,min=5"`
	Password string `json:"password" binding:"required,max=16,min=6"`
}

type LoginParam struct {
	Name     string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}
