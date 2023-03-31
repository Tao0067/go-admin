package params

type SearchAdminUser struct {
	Name string `form:"name"`
	Page int    `form:"page"`
	Size int    `form:"size"`
}

type EditAdminUser struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type DelAdminUser struct {
	Id int `json:"id"`
}
