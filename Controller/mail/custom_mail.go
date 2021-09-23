package mail

import (
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	controller "github.com/sen329/test5/Controller"
	model "github.com/sen329/test5/Model"
)

func Createcustommail(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO lokapala_accountdb.t_mail_custom_message(subject, message) VALUES (?,?)")
	if err != nil {
		panic(err.Error())
	}

	subject := r.Form.Get("subject")
	message := r.Form.Get("message")

	res, err := stmt.Exec(subject, message)
	if err != nil {
		panic(err.Error())
	}

	item_type_id := r.Form.Get("item_type_id")
	item_id := r.Form.Get("item_id")
	amount := r.Form.Get("amount")
	custom_message_id, err := res.LastInsertId()
	if err != nil {
		panic(err.Error())
	}

	if item_type_id != "" && item_id != "" && amount != "" {

		stmt2, err := db.Prepare("INSERT INTO lokapala_accountdb.t_mail_attachment(item_id, item_type, amount, custom_message_id) VALUES (?,?,?,?)")
		if err != nil {
			panic(err.Error())
		}

		_, err = stmt2.Exec(item_id, item_type_id, amount, custom_message_id)
		if err != nil {
			panic(err.Error())
		}

	}

	json.NewEncoder(w).Encode("Success")

}

func Getcustommails(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	var custom_mails []model.Custom_mail

	result, err := db.Query("SELECT * from lokapala_accountdb.t_mail_custom_message")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var custom_mail model.Custom_mail
		err := result.Scan(&custom_mail.Message_id, &custom_mail.Subject, &custom_mail.Message)
		if err != nil {
			panic(err.Error())
		}

		custom_mails = append(custom_mails, custom_mail)

	}

	json.NewEncoder(w).Encode(custom_mails)

}

func Getcustommail(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("id")
	var custom_mail model.Custom_mail

	result, err := db.Query("SELECT * from lokapala_accountdb.t_mail_custom_message where message_id = ?", id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		err := result.Scan(&custom_mail.Message_id, &custom_mail.Subject, &custom_mail.Message)
		if err != nil {
			panic(err.Error())
		}

	}

	json.NewEncoder(w).Encode(custom_mail)

}

func Updatecustommail(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("id")

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE lokapala_accountdb.t_mail_custom_message SET subject = ?, message = ? where message_id = ?")
	if err != nil {
		panic(err.Error())
	}

	subject := r.Form.Get("subject")
	message := r.Form.Get("message")

	_, err = stmt.Exec(subject, message, id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}

func Deletecustommail(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("id")

	stmt2, err := db.Prepare("DELETE FROM lokapala_accountdb.t_mail_custom_message WHERE message_id = ? AND template_id = NULL")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt2.Exec(id)
	if err != nil {
		panic(err.Error())
	}

	stmt, err := db.Prepare("DELETE FROM lokapala_accountdb.t_mail_custom_message WHERE message_id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}
