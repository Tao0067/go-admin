package params

type RegisterParam struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type LoginParam struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
