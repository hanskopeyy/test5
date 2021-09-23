package model

type Icon_frame struct {
	Frame_id     int64   `json:"frame_id"`
	Description  string  `json:"description"`
	Release_date *string `json:"release_date"`
}

type Icon_avatar struct {
	Avatar_id    int64   `json:"avatar_id"`
	Description  string  `json:"description"`
	Release_date *string `json:"release_date"`
	Free         int     `json:"free"`
}
