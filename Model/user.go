package model

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User_details struct {
	Id        string  `json:"id"`
	Name      string  `json:"name"`
	Email     string  `json:"email"`
	Role_id   *int    `josn:"role_id"`
	Role_name *string `json:"role_name"`
}
