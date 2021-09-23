package mail

import (
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	controller "github.com/sen329/test5/Controller"
	model "github.com/sen329/test5/Model"
)

func Createtemplate(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO lokapala_accountdb.t_mail_template(subject, message) VALUES (?,?)")
	if err != nil {
		panic(err.Error())
	}

	subject := r.Form.Get("subject")
	message := r.Form.Get("message")

	res, err := stmt.Exec(subject, message)
	if err != nil {
		panic(err.Error())
	}

	template_id, err := res.LastInsertId()
	if err != nil {
		panic(err.Error())
	}
	item_type_id := r.Form.Get("item_type_id")
	item_id := r.Form.Get("item_id")
	amount := r.Form.Get("amount")

	if item_type_id != "" && item_id != "" && amount != "" {

		stmt2, err := db.Prepare("INSERT INTO lokapala_accountdb.t_mail_attachment(template_id, item_id, item_type, amount) VALUES (?,?,?,?)")
		if err != nil {
			panic(err.Error())
		}

		_, err = stmt2.Exec(template_id, item_id, item_type_id, amount)
		if err != nil {
			panic(err.Error())
		}

	}

	json.NewEncoder(w).Encode("Success")

}

func Gettemplates(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	var templates []model.Mail_template

	result, err := db.Query("SELECT * from lokapala_accountdb.t_mail_template")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var template model.Mail_template
		err := result.Scan(&template.Template_id, &template.Subject, &template.Message)
		if err != nil {
			panic(err.Error())
		}

		templates = append(templates, template)

	}

	json.NewEncoder(w).Encode(templates)

}

func Gettemplate(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("id")

	var template model.Mail_template
	result, err := db.Query("SELECT * from lokapala_accountdb.t_mail_template where template_id = ?", id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {

		err := result.Scan(&template.Template_id, &template.Subject, &template.Message)
		if err != nil {
			panic(err.Error())
		}

	}

	json.NewEncoder(w).Encode(template)

}

func Updatetemplates(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("template_id")

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE lokapala_accountdb.t_mail_template SET subject = ?, message = ? where template_id = ?")
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

func DeleteTemplates(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("template_id")

	stmt2, err := db.Prepare("DELETE FROM lokapala_accountdb.t_mail_attachment WHERE template_id = ? AND custom_message_id = NULL")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt2.Exec(id)
	if err != nil {
		panic(err.Error())
	}

	stmt, err := db.Prepare("DELETE FROM lokapala_accountdb.t_mail_template WHERE template_id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}
