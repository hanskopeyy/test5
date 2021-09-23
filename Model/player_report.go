package model

type Player_report struct {
	Report_id        int64   `json:"report_id"`
	Description      string  `json:"description"`
	Room_id          int64   `json:"room_id"`
	Reporter_user_id int64   `json:"reporter_user_id"`
	Reported_user_id int64   `json:"reported_user_id"`
	Message          *string `json:"message"`
	Report_date      string  `json:"report_date"`
}

type Player_profile_report struct {
	Report_profile_id int64  `json:"report_profile_id"`
	Report_type       string `json:"report_type"`
	Reporter_user_id  int    `json:"reporter_user_id"`
	Reporter_user     string `json:"reporter_user"`
	Reported_user_id  int    `json:"reported_user_id"`
	Reported_user     string `json:"reported_user"`
	Report_date       string `json:"report_date"`
	Checked           int    `json:"checked"`
}
