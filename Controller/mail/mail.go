package mail

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	controller "github.com/sen329/test5/Controller"
	model "github.com/sen329/test5/Model"
)

func NewNullString(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

func Sendmail(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO lokapala_accountdb.t_mail(mail_type,sender_id,receiver_id,send_date,mail_template,parameter,custom_message_id) VALUES (?,?,?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	mail_type := r.Form.Get("mail_type")
	sender_id := r.Form.Get("sender_id")
	receiver_id := r.Form.Get("receiver_id")
	send_date := r.Form.Get("send_date")
	mail_template := r.Form.Get("mail_template")
	parameter := r.Form.Get("parameter")
	custom_message_id := r.Form.Get("custom_message_id")

	_, err = stmt.Exec(mail_type, NewNullString(sender_id), receiver_id, NewNullString(send_date), NewNullString(mail_template), NewNullString(parameter), NewNullString(custom_message_id))
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	fmt.Fprintf(w, "Success")
}

func Getmails(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()

	var mails []model.Mail

	result, err := db.Query("SELECT * from lokapala_accountdb.t_mail")
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	for result.Next() {
		var mail model.Mail
		err := result.Scan(&mail.Mail_id, &mail.Mail_type, &mail.Sender_id, &mail.Receiver_id, &mail.Send_date, &mail.Mail_template, &mail.Confirm_read, &mail.Read_Date, &mail.Confirm_claim, &mail.Claim_date, &mail.Parameter, &mail.Custom_message_id)
		if err != nil {
			panic(err.Error())
		}

		mails = append(mails, mail)

	}

	json.NewEncoder(w).Encode(mails)

}

func SetSendDate(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("id")

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE lokapala_accountdb.t_mail SET send_date = ? WHERE mail_id = ?")
	if err != nil {
		panic(err.Error())
	}

	send_date_new := r.Form.Get("send_date")

	_, err = stmt.Exec(send_date_new, id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}
