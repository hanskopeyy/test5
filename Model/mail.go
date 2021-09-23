package model

// type mail_type int

// const (
// 	System = iota
// 	Friend
// 	Update
// 	Gifts
// )

type Mail struct {
	Mail_id           *int    `json:"mail_id"`
	Mail_type         *string `json:"mail_type"`
	Sender_id         *int    `json:"sender_id"`
	Receiver_id       *int    `json:"reciever_id"`
	Send_date         *string `json:"send_date"`
	Mail_template     *int    `json:"mail_template"`
	Confirm_read      *int    `json:"confirm_read"`
	Read_Date         *string `json:"read_date"`
	Confirm_claim     *int    `json:"confirm_claim"`
	Claim_date        *string `json:"claim_date"`
	Parameter         *string `json:"parameter"`
	Custom_message_id *int    `json:"custom_message_id"`
}

type Mail_template struct {
	Template_id int     `json:"template_id"`
	Subject     *string `json:"subject"`
	Message     *string `json:"message"`
}

type Mail_attachment struct {
	Id                int  `json:"id"`
	Template_id       *int `json:"template_id"`
	Item_id           int  `json:"item_id"`
	Item_type         int  `json:"item_type"`
	Amount            int  `json:"amount"`
	Custom_message_id *int `json:"custom_message_id"`
}

type Mail_attachment_details struct {
	Template_id       *int    `json:"template_id"`
	Item_type_id      *int    `json:"item_type_id"`
	Item_Id           *int    `json:"item_id"`
	Item_type         *string `json:"item_type"`
	Item_name         *string `json:"item_name"`
	Amount            int     `json:"amount"`
	Custom_message_id *int    `json:"custom_message_id"`
}

type Custom_mail struct {
	Message_id int    `json:"message_id"`
	Subject    string `json:"subject"`
	Message    string `json:"message"`
}

type Login_mail struct {
	Template_id int64   `json:"template_id"`
	Parameter   *string `json:"parameter"`
	Start_date  string  `json:"start_date"`
	End_date    string  `json:"end_date"`
}
