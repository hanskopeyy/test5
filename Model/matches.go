package model

type Matches struct {
	Room_id     int64   `json:"room_id"`
	Room_name   string  `json:"room_name"`
	Match_id    int64   `json:"match_id"`
	Game_mode   *int    `json:"game_mode"`
	Server_ip   *string `json:"server_ip"`
	Server_port *int    `json:"server_port"`
	Start_time  *string `json:"start_time"`
	Can_timeout *string `json:"can_timeout"`
}
