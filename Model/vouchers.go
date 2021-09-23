package model

type Voucher struct {
	Id           int64   `json:"id"`
	Key          string  `json:"key"`
	Created_date *string `json:"created_date"`
	User_id      *int64  `json:"user_id"`
	Claimed_date *string `json:"claimed_date"`
	Expired_date *string `json:"expired_date"`
}

type Voucher_detail struct {
	Voucher_id int64  `json:"voucher_id"`
	Item_type  int64  `json:"item_type"`
	Item_id    int64  `json:"item_id"`
	Amount     int    `json:"amount"`
	Detail     string `json:"detail"`
}
