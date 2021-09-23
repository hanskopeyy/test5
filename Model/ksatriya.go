package model

type Ksatriya struct {
	Ksatriya_id   int64   `json:"ksatriya_id"`
	Role          *string `json:"role"`
	Release_date  *string `json:"release_date"`
	Ksatriya_name *string `json:"ksatriya_name"`
}

type Ksatriya_fragment struct {
	Ksatriya_id      int64 `json:"ksatriya_id"`
	Amount_needed    int   `json:"amount"`
	Sell_currency_id int   `json:"sell_currency_id"`
	Sell_value       int   `json:"sell_value"`
}

type Ksatriya_skin_part struct {
	Skin_part_id int    `json:"skin_part_id"`
	Release_date string `json:"release_date"`
}
