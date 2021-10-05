package controller

import (
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	model "github.com/sen329/test5/Model"
)

func GetAllPlayerReports(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()

	count := r.URL.Query().Get("count")
	offset := r.URL.Query().Get("offset")

	var reports []model.Player_report

	query, err := db.Prepare(`SELECT report_id, tt.description as report_description, t.room_id, reporter_user_id, uu.user_name, reported_user_id, u.user_name, COALESCE(message, '')  as message, report_date, checked FROM lokapala_playerreportdb.t_player_report t LEFT JOIN lokapala_playerreportdb.t_player_report_type tt on t.report_type = tt.report_type_id LEFT JOIN lokapala_accountdb.t_user u ON t.reported_user_id = u.user_id LEFT JOIN lokapala_accountdb.t_user uu ON t.reporter_user_id = uu.user_id LIMIT ? OFFSET ?`)
	if err != nil {
		panic(err.Error())
	}

	result, err := query.Query(count, offset)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var report model.Player_report
		err := result.Scan(&report.Report_id, &report.Description, &report.Room_id, &report.Reporter_user_id, &report.Reporter_user, &report.Reported_user_id, &report.Reported_user, &report.Message, &report.Report_date, &report.Checked)
		if err != nil {
			panic(err.Error())
		}

		reports = append(reports, report)
	}

	json.NewEncoder(w).Encode(reports)

}

func GetAllPlayerReportsByReportedUser(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()

	count := r.URL.Query().Get("count")
	offset := r.URL.Query().Get("offset")
	user_id := r.URL.Query().Get("user_id")

	var reports []model.Player_report

	query, err := db.Prepare(`SELECT report_id, tt.description as report_description, t.room_id, reporter_user_id, uu.user_name, reported_user_id, u.user_name, COALESCE(message, '')  as message, report_date, checked FROM lokapala_playerreportdb.t_player_report t LEFT JOIN lokapala_playerreportdb.t_player_report_type tt on t.report_type = tt.report_type_id LEFT JOIN lokapala_accountdb.t_user u ON t.reported_user_id = u.user_id LEFT JOIN lokapala_accountdb.t_user uu ON t.reporter_user_id = uu.user_id WHERE t.reported_user_id = ? LIMIT ? OFFSET ? `)
	if err != nil {
		panic(err.Error())
	}

	result, err := query.Query(user_id, count, offset)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var report model.Player_report
		err := result.Scan(&report.Report_id, &report.Description, &report.Room_id, &report.Reporter_user_id, &report.Reporter_user, &report.Reported_user_id, &report.Reported_user, &report.Message, &report.Report_date, &report.Checked)
		if err != nil {
			panic(err.Error())
		}

		reports = append(reports, report)
	}

	json.NewEncoder(w).Encode(reports)

}

func GetAllPlayerReportsByReporterUser(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()

	count := r.URL.Query().Get("count")
	offset := r.URL.Query().Get("offset")
	user_id := r.URL.Query().Get("user_id")

	var reports []model.Player_report

	query, err := db.Prepare("SELECT report_id, tt.description as report_description, t.room_id, reporter_user_id, uu.user_name, reported_user_id, u.user_name, COALESCE(message, '')  as message, report_date, checked FROM lokapala_playerreportdb.t_player_report t LEFT JOIN lokapala_playerreportdb.t_player_report_type tt on t.report_type = tt.report_type_id LEFT JOIN lokapala_accountdb.t_user u ON t.reported_user_id = u.user_id LEFT JOIN lokapala_accountdb.t_user uu ON t.reporter_user_id = uu.user_id WHERE t.reporter_user_id = ? LIMIT ? OFFSET ?")
	if err != nil {
		panic(err.Error())
	}

	result, err := query.Query(user_id, count, offset)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var report model.Player_report
		err := result.Scan(&report.Report_id, &report.Description, &report.Room_id, &report.Reporter_user_id, &report.Reporter_user, &report.Reported_user_id, &report.Reported_user, &report.Message, &report.Report_date, &report.Checked)
		if err != nil {
			panic(err.Error())
		}

		reports = append(reports, report)
	}

	json.NewEncoder(w).Encode(reports)

}

func GetAllPlayerReportsByRoom(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()

	count := r.URL.Query().Get("count")
	offset := r.URL.Query().Get("offset")
	room_id := r.URL.Query().Get("room_id")

	var reports []model.Player_report

	query, err := db.Prepare("SELECT report_id, tt.description as report_description, t.room_id, reporter_user_id, uu.user_name, reported_user_id, u.user_name, COALESCE(message, '')  as message, report_date, checked FROM lokapala_playerreportdb.t_player_report t LEFT JOIN lokapala_playerreportdb.t_player_report_type tt on t.report_type = tt.report_type_id LEFT JOIN lokapala_accountdb.t_user u ON t.reported_user_id = u.user_id LEFT JOIN lokapala_accountdb.t_user uu ON t.reporter_user_id = uu.user_id WHERE t.room_id = ? LIMIT ? OFFSET ?")
	if err != nil {
		panic(err.Error())
	}

	result, err := query.Query(room_id, count, offset)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var report model.Player_report
		err := result.Scan(&report.Report_id, &report.Description, &report.Room_id, &report.Reporter_user_id, &report.Reporter_user, &report.Reported_user_id, &report.Reported_user, &report.Message, &report.Report_date, &report.Checked)
		if err != nil {
			panic(err.Error())
		}

		reports = append(reports, report)
	}

	json.NewEncoder(w).Encode(reports)

}

func GetPlayerReport(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()

	report_id := r.URL.Query().Get("report_id")

	var reports []model.Player_report

	query, err := db.Prepare(`SELECT report_id, tt.description as report_description, t.room_id, reporter_user_id, uu.user_name, reported_user_id, u.user_name, COALESCE(message, '')  as message, report_date, checked FROM lokapala_playerreportdb.t_player_report t LEFT JOIN lokapala_playerreportdb.t_player_report_type tt on t.report_type = tt.report_type_id LEFT JOIN lokapala_accountdb.t_user u ON t.reported_user_id = u.user_id LEFT JOIN lokapala_accountdb.t_user uu ON t.reporter_user_id = uu.user_id WHERE t.report_id = ?`)
	if err != nil {
		panic(err.Error())
	}

	result, err := query.Query(report_id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var report model.Player_report
		err := result.Scan(&report.Report_id, &report.Description, &report.Room_id, &report.Reporter_user_id, &report.Reporter_user, &report.Reported_user_id, &report.Reported_user, &report.Message, &report.Report_date, &report.Checked)
		if err != nil {
			panic(err.Error())
		}

		reports = append(reports, report)
	}

	json.NewEncoder(w).Encode(reports)

}

func ApprovePlayerReport(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()

	report_id := r.URL.Query().Get("report_id")

	stmt, err := db.Prepare(`UPDATE lokapala_playerreportdb.t_player_report SET checked = 1 WHERE report_id = ?`)
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(report_id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}

func GetAllPlayerProfileReports(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()

	var reports []model.Player_profile_report

	result, err := db.Query("SELECT A.report_profile_id as report_profile_id, B.description as report_type,C.user_id as reporter_id, C.user_name as reporter_user,D.user_id as reported_user_id, D.user_name as reported_user, A.report_date, A.checked FROM lokapala_playerreportdb.t_player_report_profile A LEFT JOIN lokapala_playerreportdb.t_player_report_type B ON B.report_type_id = A.report_type LEFT JOIN lokapala_accountdb.t_user C ON C.user_id = A.reporter_user_id LEFT JOIN lokapala_accountdb.t_user D ON D.user_id = A.reported_user_id")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var report model.Player_profile_report
		err := result.Scan(&report.Report_profile_id, &report.Report_type, &report.Reporter_user_id, &report.Reporter_user, &report.Reported_user_id, &report.Reported_user, &report.Report_date, &report.Checked)
		if err != nil {
			panic(err.Error())
		}

		reports = append(reports, report)
	}

	json.NewEncoder(w).Encode(reports)

}

func GetAllPlayerProfileReportsByReporterUser(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()

	user_id := r.URL.Query().Get("user_id")

	var reports []model.Player_profile_report

	query, err := db.Prepare("SELECT A.report_profile_id as report_profile_id, B.description as report_type,C.user_id as reporter_id, C.user_name as reporter_user,D.user_id as reported_user_id, D.user_name as reported_user, A.report_date, A.checked FROM lokapala_playerreportdb.t_player_report_profile A LEFT JOIN lokapala_playerreportdb.t_player_report_type B ON B.report_type_id = A.report_type LEFT JOIN lokapala_accountdb.t_user C ON C.user_id = A.reporter_user_id LEFT JOIN lokapala_accountdb.t_user D ON D.user_id = A.reported_user_id WHERE A.reporter_user_id = ?")
	if err != nil {
		panic(err.Error())
	}

	result, err := query.Query(user_id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var report model.Player_profile_report
		err := result.Scan(&report.Report_profile_id, &report.Report_type, &report.Reporter_user_id, &report.Reporter_user, &report.Reported_user_id, &report.Reported_user, &report.Report_date, &report.Checked)
		if err != nil {
			panic(err.Error())
		}

		reports = append(reports, report)
	}

	json.NewEncoder(w).Encode(reports)

}

func GetAllPlayerProfileReportsByReportedUser(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()

	user_id := r.URL.Query().Get("user_id")

	var reports []model.Player_profile_report

	query, err := db.Prepare("SELECT A.report_profile_id as report_profile_id, B.description as report_type,C.user_id as reporter_id, C.user_name as reporter_user, D.user_name as reported_user, A.report_date, A.checked FROM lokapala_playerreportdb.t_player_report_profile A LEFT JOIN lokapala_playerreportdb.t_player_report_type B ON B.report_type_id = A.report_type LEFT JOIN lokapala_accountdb.t_user C ON C.user_id = A.reporter_user_id LEFT JOIN lokapala_accountdb.t_user D ON D.user_id = A.reported_user_id WHERE A.reported_user_id = ?")
	if err != nil {
		panic(err.Error())
	}

	result, err := query.Query(user_id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var report model.Player_profile_report
		err := result.Scan(&report.Report_profile_id, &report.Report_type, &report.Reporter_user_id, &report.Reporter_user, &report.Reported_user, &report.Report_date, &report.Checked)
		if err != nil {
			panic(err.Error())
		}

		reports = append(reports, report)
	}

	json.NewEncoder(w).Encode(reports)

}

func GetPlayerProfileReport(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()

	player_report_id := r.URL.Query().Get("player_report_id")

	query, err := db.Prepare("SELECT A.report_profile_id as report_profile_id, B.description as report_type,C.user_id as reporter_id, C.user_name as reporter_user,D.user_id as reported_user_id, D.user_name as reported_user, A.report_date, A.checked FROM lokapala_playerreportdb.t_player_report_profile A LEFT JOIN lokapala_playerreportdb.t_player_report_type B ON B.report_type_id = A.report_type LEFT JOIN lokapala_accountdb.t_user C ON C.user_id = A.reporter_user_id LEFT JOIN lokapala_accountdb.t_user D ON D.user_id = A.reported_user_id WHERE A.report_profile_id = ?")
	if err != nil {
		panic(err.Error())
	}

	result, err := query.Query(player_report_id)
	if err != nil {
		panic(err.Error())
	}

	var report model.Player_profile_report

	for result.Next() {

		err := result.Scan(&report.Report_profile_id, &report.Report_type, &report.Reporter_user_id, &report.Reporter_user, &report.Reported_user_id, &report.Reported_user, &report.Report_date, &report.Checked)
		if err != nil {
			panic(err.Error())
		}

	}

	json.NewEncoder(w).Encode(report)

}

func ApprovePlayerProfileReport(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()

	report_id := r.URL.Query().Get("report_id")

	stmt, err := db.Prepare(`UPDATE lokapala_playerreportdb.t_player_report_profile SET checked = 1 WHERE report_profile_id = ?`)
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(report_id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}
