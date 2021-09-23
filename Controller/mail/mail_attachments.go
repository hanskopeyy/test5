package mail

import (
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	controller "github.com/sen329/test5/Controller"
	model "github.com/sen329/test5/Model"
)

func Attachitem(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO lokapala_accountdb.t_mail_attachment(template_id, item_id, item_type, amount, custom_message_id) VALUES (?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	template_id := r.Form.Get("template_id")
	item_id := r.Form.Get("item_id")
	item_type := r.Form.Get("item_type")
	amount := r.Form.Get("amount")
	custom_message_id := r.Form.Get("custom_message_id")

	_, err = stmt.Exec(NewNullString(template_id), item_id, item_type, amount, NewNullString(custom_message_id))
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}

func Getmailattachments(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	var attachments []model.Mail_attachment_details

	result, err := db.Query("SELECT A.template_id,B.item_type_name,A.item_type,A.item_id, CASE WHEN A.item_type = 1 THEN (SELECT name FROM lokapala_accountdb.t_currency_type curr WHERE curr.currency_id = A.item_id ) WHEN A.item_type = 2 THEN (SELECT ksatriya_name FROM lokapala_accountdb.t_ksatriya ksa WHERE ksa.ksatriya_id = A.item_id) WHEN A.item_type = 3 THEN (SELECT CONCAT_WS(" + `" - "` + ", ksa_skin.ksatriya_skin_id, ksa.ksatriya_name ) FROM lokapala_accountdb.t_ksatriya_skin ksa_skin LEFT JOIN lokapala_accountdb.t_ksatriya ksa ON ksa_skin.ksatriya_id = ksa.ksatriya_id WHERE ksa_skin.ksatriya_skin_id = A.item_id) WHEN A.item_type = 4 THEN (SELECT rune.name FROM lokapala_accountdb.t_rune as rune WHERE rune.rune_id = A.item_id) WHEN A.item_type = 5 THEN (SELECT item.misc_name FROM lokapala_accountdb.t_misc_item item WHERE item.misc_id = A.item_id) WHEN A.item_type = 6 THEN (SELECT box.box_name FROM lokapala_accountdb.t_box box WHERE box.box_id = A.item_id) WHEN A.item_type = 7 THEN (SELECT chest.duration FROM lokapala_accountdb.t_chest chest WHERE chest.duration = A.item_id) WHEN A.item_type = 8 THEN (SELECT energy.description FROM lokapala_accountdb.t_energy energy WHERE energy_id = A.item_id) WHEN A.item_type = 9 THEN (SELECT skin_part.skin_part_id FROM lokapala_accountdb.t_ksatriya_skin_part skin_part WHERE skin_part_id = A.item_id) WHEN A.item_type = 10 THEN (SELECT CONCAT_WS(" + `" - "` + ",premium.item_id, premium.duration) FROM lokapala_accountdb.t_premium premium WHERE premium.item_id) WHEN A.item_type = 11 THEN (SELECT frame.description FROM lokapala_accountdb.t_icon_frame frame WHERE frame.frame_id = A.item_id) WHEN A.item_type = 12 THEN (SELECT avatar.description FROM lokapala_accountdb.t_icon_avatar avatar WHERE avatar.avatar_id = A.item_id) WHEN A.item_type = 14 THEN (SELECT vahana.vahana_skin FROM lokapala_accountdb.t_vahana_skin vahana WHERE vahana.vahana_skin_id = A.item_id) WHEN A.item_type = 15 THEN (SELECT ksa_frag.ksatriya_id FROM lokapala_accountdb.t_ksatriya_fragment ksa_frag WHERE ksa_frag.ksatriya_id = A.item_id) WHEN A.item_type = 16 THEN (SELECT ksa_skin_frag.ksatriya_skin_id FROM lokapala_accountdb.t_ksatriya_skin_fragment ksa_skin_frag WHERE ksa_skin_frag.ksatriya_skin_id = A.item_id) END AS item_name, A.amount, A.custom_message_id FROM lokapala_accountdb.t_mail_attachment A LEFT JOIN lokapala_accountdb.t_item_type B ON A.item_type = B.item_type_id")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var attachment model.Mail_attachment_details
		err := result.Scan(&attachment.Template_id, &attachment.Item_type, &attachment.Item_type_id, &attachment.Item_Id, &attachment.Item_name, &attachment.Amount, &attachment.Custom_message_id)
		if err != nil {
			panic(err.Error())
		}

		attachments = append(attachments, attachment)

	}

	json.NewEncoder(w).Encode(attachments)

}

func GetmailattachmentByItem(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	item_type := r.URL.Query().Get("item_type")
	item_id := r.URL.Query().Get("item_id")
	var attachment model.Mail_attachment

	query, err := db.Prepare("SELECT * from lokapala_accountdb.t_mail_attachment WHERE item_type = ? AND item_id = ? AND template_id = NULL AND custom_message_id = NULL")
	if err != nil {
		panic(err.Error())
	}

	result, err := query.Query(item_type, item_id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		err := result.Scan(&attachment.Template_id, &attachment.Item_id, &attachment.Item_type, &attachment.Amount, &attachment.Custom_message_id)
		if err != nil {
			panic(err.Error())
		}

	}

	json.NewEncoder(w).Encode(attachment)
}

func GetmailattachmentByTemplateId(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	template_id := r.URL.Query().Get("template_id")
	item_type := r.URL.Query().Get("item_type")
	item_id := r.URL.Query().Get("item_id")
	var attachment model.Mail_attachment

	query, err := db.Prepare("SELECT * from lokapala_accountdb.t_mail_attachment where template_id = ? AND item_type = ? AND item_id = ?")
	if err != nil {
		panic(err.Error())
	}

	result, err := query.Query(template_id, item_type, item_id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		err := result.Scan(&attachment.Template_id, &attachment.Item_id, &attachment.Item_type, &attachment.Amount, &attachment.Custom_message_id)
		if err != nil {
			panic(err.Error())
		}

	}

	json.NewEncoder(w).Encode(attachment)
}

func GetmailattachmentByTemplateIdOnly(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	template_id := r.URL.Query().Get("template_id")
	var attachments []model.Mail_attachment_details

	query, err := db.Prepare("SELECT A.template_id,B.item_type_name , A.item_type,A.item_id, CASE WHEN A.item_type = 1 THEN (SELECT name FROM lokapala_accountdb.t_currency_type curr WHERE curr.currency_id = A.item_id ) WHEN A.item_type = 2 THEN (SELECT ksatriya_name FROM lokapala_accountdb.t_ksatriya ksa WHERE ksa.ksatriya_id = A.item_id) WHEN A.item_type = 3 THEN (SELECT CONCAT_WS(" + `" - "` + ", ksa_skin.ksatriya_skin_id, ksa.ksatriya_name ) FROM lokapala_accountdb.t_ksatriya_skin ksa_skin LEFT JOIN lokapala_accountdb.t_ksatriya ksa ON ksa_skin.ksatriya_id = ksa.ksatriya_id WHERE ksa_skin.ksatriya_skin_id = A.item_id) WHEN A.item_type = 4 THEN (SELECT rune.name FROM lokapala_accountdb.t_rune as rune WHERE rune.rune_id = A.item_id) WHEN A.item_type = 5 THEN (SELECT item.misc_name FROM lokapala_accountdb.t_misc_item item WHERE item.misc_id = A.item_id) WHEN A.item_type = 6 THEN (SELECT box.box_name FROM lokapala_accountdb.t_box box WHERE box.box_id = A.item_id) WHEN A.item_type = 7 THEN (SELECT chest.duration FROM lokapala_accountdb.t_chest chest WHERE chest.duration = A.item_id) WHEN A.item_type = 8 THEN (SELECT energy.description FROM lokapala_accountdb.t_energy energy WHERE energy_id = A.item_id) WHEN A.item_type = 9 THEN (SELECT skin_part.skin_part_id FROM lokapala_accountdb.t_ksatriya_skin_part skin_part WHERE skin_part_id = A.item_id) WHEN A.item_type = 10 THEN (SELECT CONCAT_WS(" + `" - "` + ",premium.item_id, premium.duration) FROM lokapala_accountdb.t_premium premium WHERE premium.item_id) WHEN A.item_type = 11 THEN (SELECT frame.description FROM lokapala_accountdb.t_icon_frame frame WHERE frame.frame_id = A.item_id) WHEN A.item_type = 12 THEN (SELECT avatar.description FROM lokapala_accountdb.t_icon_avatar avatar WHERE avatar.avatar_id = A.item_id) WHEN A.item_type = 14 THEN (SELECT vahana.vahana_skin FROM lokapala_accountdb.t_vahana_skin vahana WHERE vahana.vahana_skin_id = A.item_id) WHEN A.item_type = 15 THEN (SELECT ksa_frag.ksatriya_id FROM lokapala_accountdb.t_ksatriya_fragment ksa_frag WHERE ksa_frag.ksatriya_id = A.item_id) WHEN A.item_type = 16 THEN (SELECT ksa_skin_frag.ksatriya_skin_id FROM lokapala_accountdb.t_ksatriya_skin_fragment ksa_skin_frag WHERE ksa_skin_frag.ksatriya_skin_id = A.item_id) END AS item_name, A.amount, A.custom_message_id FROM lokapala_accountdb.t_mail_attachment A LEFT JOIN lokapala_accountdb.t_item_type B ON A.item_type = B.item_type_id WHERE A.template_id = ?")
	if err != nil {
		panic(err.Error())
	}

	result, err := query.Query(template_id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var attachment model.Mail_attachment_details
		err := result.Scan(&attachment.Template_id, &attachment.Item_type, &attachment.Item_type_id, &attachment.Item_Id, &attachment.Item_name, &attachment.Amount, &attachment.Custom_message_id)
		if err != nil {
			panic(err.Error())
		}

		attachments = append(attachments, attachment)

	}

	json.NewEncoder(w).Encode(attachments)
}

func GetmailattachmentByCustomMessageId(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	custom_id := r.URL.Query().Get("custom_message_id")
	item_type := r.URL.Query().Get("item_type")
	item_id := r.URL.Query().Get("item_id")
	var attachment model.Mail_attachment

	query, err := db.Prepare("SELECT * from lokapala_accountdb.t_mail_attachment where custom_message_id = ? AND item_type = ? AND item_id = ?")
	if err != nil {
		panic(err.Error())
	}

	result, err := query.Query(custom_id, item_type, item_id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		err := result.Scan(&attachment.Template_id, &attachment.Item_id, &attachment.Item_type, &attachment.Amount, &attachment.Custom_message_id)
		if err != nil {
			panic(err.Error())
		}

	}

	json.NewEncoder(w).Encode(attachment)
}

func GetmailattachmentByCustomMessageIdOnly(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	custom_id := r.URL.Query().Get("custom_message_id")
	var attachments []model.Mail_attachment_details

	query, err := db.Prepare("SELECT A.template_id,B.item_type_name , A.item_type, A.item_id, CASE WHEN A.item_type = 1 THEN (SELECT name FROM lokapala_accountdb.t_currency_type curr WHERE curr.currency_id = A.item_id ) WHEN A.item_type = 2 THEN (SELECT ksatriya_name FROM lokapala_accountdb.t_ksatriya ksa WHERE ksa.ksatriya_id = A.item_id) WHEN A.item_type = 3 THEN (SELECT CONCAT_WS(" + `" - "` + ", ksa_skin.ksatriya_skin_id, ksa.ksatriya_name ) FROM lokapala_accountdb.t_ksatriya_skin ksa_skin LEFT JOIN lokapala_accountdb.t_ksatriya ksa ON ksa_skin.ksatriya_id = ksa.ksatriya_id WHERE ksa_skin.ksatriya_skin_id = A.item_id) WHEN A.item_type = 4 THEN (SELECT rune.name FROM lokapala_accountdb.t_rune as rune WHERE rune.rune_id = A.item_id) WHEN A.item_type = 5 THEN (SELECT item.misc_name FROM lokapala_accountdb.t_misc_item item WHERE item.misc_id = A.item_id) WHEN A.item_type = 6 THEN (SELECT box.box_name FROM lokapala_accountdb.t_box box WHERE box.box_id = A.item_id) WHEN A.item_type = 7 THEN (SELECT chest.duration FROM lokapala_accountdb.t_chest chest WHERE chest.duration = A.item_id) WHEN A.item_type = 8 THEN (SELECT energy.description FROM lokapala_accountdb.t_energy energy WHERE energy_id = A.item_id) WHEN A.item_type = 9 THEN (SELECT skin_part.skin_part_id FROM lokapala_accountdb.t_ksatriya_skin_part skin_part WHERE skin_part_id = A.item_id) WHEN A.item_type = 10 THEN (SELECT CONCAT_WS(" + `" - "` + ",premium.item_id, premium.duration) FROM lokapala_accountdb.t_premium premium WHERE premium.item_id) WHEN A.item_type = 11 THEN (SELECT frame.description FROM lokapala_accountdb.t_icon_frame frame WHERE frame.frame_id = A.item_id) WHEN A.item_type = 12 THEN (SELECT avatar.description FROM lokapala_accountdb.t_icon_avatar avatar WHERE avatar.avatar_id = A.item_id) WHEN A.item_type = 14 THEN (SELECT vahana.vahana_skin FROM lokapala_accountdb.t_vahana_skin vahana WHERE vahana.vahana_skin_id = A.item_id) WHEN A.item_type = 15 THEN (SELECT ksa_frag.ksatriya_id FROM lokapala_accountdb.t_ksatriya_fragment ksa_frag WHERE ksa_frag.ksatriya_id = A.item_id) WHEN A.item_type = 16 THEN (SELECT ksa_skin_frag.ksatriya_skin_id FROM lokapala_accountdb.t_ksatriya_skin_fragment ksa_skin_frag WHERE ksa_skin_frag.ksatriya_skin_id = A.item_id) END AS item_name, A.amount, A.custom_message_id FROM lokapala_accountdb.t_mail_attachment A LEFT JOIN lokapala_accountdb.t_item_type B ON A.item_type = B.item_type_id WHERE A.custom_message_id = ?")
	if err != nil {
		panic(err.Error())
	}

	result, err := query.Query(custom_id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var attachment model.Mail_attachment_details
		err := result.Scan(&attachment.Template_id, &attachment.Item_type, &attachment.Item_type_id, &attachment.Item_Id, &attachment.Item_name, &attachment.Amount, &attachment.Custom_message_id)
		if err != nil {
			panic(err.Error())
		}
		attachments = append(attachments, attachment)
	}

	json.NewEncoder(w).Encode(attachments)
}

func UpdatemailattachmentByItem(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	item_type_old := r.URL.Query().Get("item_type")
	item_id_old := r.URL.Query().Get("item_id")
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE lokapala_accountdb.t_mail_attachment SET template_id, item_id = ?, item_type = ?, amount = ?, custom_message_id WHERE item_type = ? AND item_id =? AND template_id = NULL AND custom_message_id = NULL")
	if err != nil {
		panic(err.Error())
	}

	template_id := r.Form.Get("template_id")
	item_id_new := r.Form.Get("item_id")
	item_type_new := r.Form.Get("item_type")
	amount := r.Form.Get("amount")
	custom_message_id := r.Form.Get("custom_message_id")

	_, err = stmt.Exec(template_id, item_id_new, item_type_new, amount, custom_message_id, item_type_old, item_id_old)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func UpdatemailattachmentByTemplateId(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	template_id := r.URL.Query().Get("template_id")
	item_type_old := r.URL.Query().Get("item_type")
	item_id_old := r.URL.Query().Get("item_id")
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE lokapala_accountdb.t_mail_attachment SET item_id = ?, item_type = ?, amount = ? WHERE template_id = ? AND item_type = ? AND item_id =?")
	if err != nil {
		panic(err.Error())
	}

	item_id_new := r.Form.Get("item_id")
	item_type_new := r.Form.Get("item_type")
	amount := r.Form.Get("amount")

	_, err = stmt.Exec(item_id_new, item_type_new, amount, template_id, item_type_old, item_id_old)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func UpdatemailattachmentByCustomMessageId(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	custom_message_id := r.URL.Query().Get("custom_message_id")
	item_type_old := r.URL.Query().Get("item_type")
	item_id_old := r.URL.Query().Get("item_id")
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE lokapala_accountdb.t_mail_attachment SET item_id = ?, item_type = ?, amount = ? WHERE custom_message_id = ? AND item_type = ? AND item_id = ?")
	if err != nil {
		panic(err.Error())
	}

	item_id := r.Form.Get("item_id")
	item_type_new := r.Form.Get("item_type")
	amount := r.Form.Get("amount")

	_, err = stmt.Exec(item_id, item_type_new, amount, custom_message_id, item_type_old, item_id_old)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func RemoveitemByItem(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	item_id := r.URL.Query().Get("item_id")
	item_type := r.URL.Query().Get("item_type")

	stmt, err := db.Prepare("DELETE FROM lokapala_accountdb.t_mail_attachment WHERE template_id = ? AND item_id = ? AND item_type = ? AND template_id = NULL AND custom_message_id = NULL")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(item_id, item_type)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}

func RemoveitemByTemplateId(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	template_id := r.URL.Query().Get("template_id")
	item_id := r.URL.Query().Get("item_id")
	item_type := r.URL.Query().Get("item_type")

	stmt, err := db.Prepare("DELETE FROM lokapala_accountdb.t_mail_attachment WHERE template_id = ? AND item_id = ? AND item_type = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(template_id, item_id, item_type)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}

func RemoveitemByCustomMessageId(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	custom_message_id := r.URL.Query().Get("custom_message_id")
	item_id := r.URL.Query().Get("item_id")
	item_type := r.URL.Query().Get("item_type")

	stmt, err := db.Prepare("DELETE FROM lokapala_accountdb.t_mail_attachment WHERE custom_message_id = ? AND item_id = ? AND item_type = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(custom_message_id, item_id, item_type)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}
