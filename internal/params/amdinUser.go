package params

type SearchAdminUser struct {
	Name string `form:"name"`
	Page int    `form:"page"`
	Size int    `form:"size"`
}
