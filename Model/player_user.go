package model

type Player struct {
	User_id      int64   `json:"user_id"`
	User_name    string  `json:"user_name"`
	Avatar_Icon  int     `json:"avatar_icon"`
	Karma        int     `json:"karma"`
	Gender       *string `json:"gender"`
	Country      int     `json:"country"`
	Role         int     `json:"role"`
	Playing_time int     `json:"playing_time"`
	Frame        int     `json:"frame"`
	Referal_id   *string `json:"referal_id"`
}

type Player_Ksatriya_ranking struct {
	User_id      int64 `json:"user_id"`
	Ksatriya_id  int64 `json:"ksatriya_id"`
	Win_count    int   `json:"win_count"`
	Lose_count   int   `json:"lose_count"`
	Match_count  int   `json:"match_count"`
	Win_rate     int   `json:"win_rate"`
	Rank         int   `json:"rank"`
	Country_rank int   `json:"country_rank"`
}

type Player_match_history struct {
	Room_id               int64  `json:"room_id"`
	User_id               *int64 `json:"user_id"`
	Win                   int    `json:"win"`
	Ksatriya_id           int64  `json:"ksatriya_id"`
	Level                 int    `json:"level"`
	Kill                  int    `json:"kill"`
	Death                 int    `json:"death"`
	Assist                int    `json:"assist"`
	Gold                  int    `json:"gold"`
	Damage_dealt          int    `json:"damage_dealt"`
	Damage_taken          int    `json:"damage_taken"`
	Ksatriya_damage_dealt int    `json:"ksatriya_damage_dealt"`
	Game_duration         int    `json:"game_duration"`
	Game_mode             *int   `json:"game_mode"`
}
