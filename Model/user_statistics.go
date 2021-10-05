package model

type Daily_user_unique struct {
	Count int `json:"count"`
}

type Daily_user struct {
	Count int `json:"count"`
	Day   int `json:"day"`
	Month int `json:"month"`
	Year  int `json:"year"`
}

type Concurrent_user struct {
	Count  int `json:"count"`
	Minute int `json:"minute"`
	Hour   int `json:"hour"`
	Day    int `json:"day"`
	Month  int `json:"month"`
	Year   int `json:"year"`
}

type User_login_type struct {
	Count        int    `json:"count"`
	Account_type string `json:"account_type"`
}

type Ksa_stats struct {
	Ksatriya_id int `json:"ksatriya_id"`
	Match_count int `json:"match_count"`
	Win_count   int `json:"win_count"`
	Lose_count  int `json:"lose_count"`
}

type Most_ksa_owned struct {
	Ksatriya_id   int    `json:"ksatriya_id"`
	Ksatriya_name string `json:"ksatriya_name"`
	Player_owned  int    `json:"player_owned"`
}

type Ksa_kda_stats struct {
	Ksatriya_id     int     `json:"ksatriya_id"`
	Ksatriya_name   string  `json:"ksatriya_name"`
	Kill_count      int     `json:"kill_count"`
	Death_count     int     `json:"death_count"`
	Assist_count    int     `json:"assist_count"`
	Kill_death_rate float64 `json:"kill_death_rate"`
}
