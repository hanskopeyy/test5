package model

type Roles struct {
	Id          string `json:"id"`
	Role_name   string `json:"role_name"`
	Description string `json:"description"`
	Active      int    `json:"active"`
}

type Users_roles struct {
	User_id   int    `json:"user_id"`
	User_name string `json:"user_name"`
	Role_id   int    `json:"role_id"`
	Role_name string `json:"role_name"`
}
