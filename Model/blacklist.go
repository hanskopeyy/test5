package model

type Blacklist struct {
	Blacklist_user_id int64  `json:"blacklist_user_id"`
	Target_user_id    int64  `json:"target_user_id"`
	Blacklist_date    string `json:"blacklist_date"`
}
