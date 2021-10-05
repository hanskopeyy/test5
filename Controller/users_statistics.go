package controller

import (
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	model "github.com/sen329/test5/Model"
)

func GetDailyUserCountUnique(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()

	start_date := r.URL.Query().Get("start_date")
	end_date := r.URL.Query().Get("end_date")

	query, err := db.Prepare("call lokapala_admindb.p_daily_active_user_unique(?,?)")
	if err != nil {
		panic(err.Error())
	}

	result, err := query.Query(start_date, end_date)
	if err != nil {
		panic(err.Error())
	}

	var stat model.Daily_user_unique

	for result.Next() {

		err := result.Scan(&stat.Count)
		if err != nil {
			panic(err.Error())
		}

	}

	json.NewEncoder(w).Encode(stat)

}

func GetDailyUserCount(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	var stats []model.Daily_user

	start_date := r.URL.Query().Get("start_date")
	end_date := r.URL.Query().Get("end_date")

	query, err := db.Prepare("call lokapala_admindb.p_daily_active_user(?,?)")
	if err != nil {
		panic(err.Error())
	}

	result, err := query.Query(start_date, end_date)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var stat model.Daily_user

		err := result.Scan(&stat.Count, &stat.Day, &stat.Month, &stat.Year)
		if err != nil {
			panic(err.Error())
		}

		stats = append(stats, stat)

	}

	json.NewEncoder(w).Encode(stats)

}

func GetConcurrentUserCount(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()

	var stats []model.Concurrent_user

	start_date := r.URL.Query().Get("start_date")
	end_date := r.URL.Query().Get("end_date")

	query, err := db.Prepare("call lokapala_admindb.p_ccu(?,?)")
	if err != nil {
		panic(err.Error())
	}

	result, err := query.Query(start_date, end_date)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var stat model.Concurrent_user
		err := result.Scan(&stat.Count, &stat.Minute, &stat.Hour, &stat.Day, &stat.Month, &stat.Year)
		if err != nil {
			panic(err.Error())
		}

		stats = append(stats, stat)

	}

	json.NewEncoder(w).Encode(stats)

}

func GetUserLoginTypeCount(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()

	var stats []model.User_login_type

	result, err := db.Query("SELECT * FROM lokapala_admindb.v_registered_user")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var stat model.User_login_type
		err := result.Scan(&stat.Count, &stat.Account_type)
		if err != nil {
			panic(err.Error())
		}

		stats = append(stats, stat)

	}

	json.NewEncoder(w).Encode(stats)

}

func GetKsaStatCount(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()

	game_mode := r.URL.Query().Get("game_mode")

	var stats []model.Ksa_stats

	result, err := db.Query("SELECT t.ksatriya_id, COUNT(1) as match_count, COUNT(case when t.win = 1 then 1 end) as win_count, COUNT(1) - COUNT(case when t.win = 1 then 1 end) as lose_count FROM (SELECT trrs.room_id, trrs.ksatriya_id, trrs.win FROM lokapala_roomdb.t_room_result_slot trrs) t JOIN lokapala_roomdb.t_past_room r ON t.room_id = r.room_id JOIN lokapala_roomdb.t_room_result rr ON t.room_id = rr.room_id WHERE COALESCE(rr.match_id, 0) > 0 AND r.start_time >= '2020-05-19 17:00:00' AND IF(? = 0, TRUE, rr.game_mode = ?) AND NOT t.ksatriya_id = 901 GROUP BY t.ksatriya_id ORDER BY match_count DESC;", game_mode, game_mode)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var stat model.Ksa_stats
		err := result.Scan(&stat.Ksatriya_id, &stat.Match_count, &stat.Win_count, &stat.Lose_count)
		if err != nil {
			panic(err.Error())
		}

		stats = append(stats, stat)

	}

	json.NewEncoder(w).Encode(stats)

}

func GetKsaTotalOwned(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()

	var stats []model.Most_ksa_owned

	result, err := db.Query("SELECT x.ksatriya_id,x.ksatriya_name, COUNT(x.ksatriya_id) AS player_owned FROM (SELECT a.ksatriya_id, b.ksatriya_name FROM lokapala_accountdb.t_inventory_ksatriya a LEFT JOIN lokapala_accountdb.t_ksatriya b ON a.ksatriya_id = b.ksatriya_id UNION ALL SELECT c.ksatriya_id, d.ksatriya_name FROM lokapala_accountdb.t_inventory_ksatriya_trial c LEFT JOIN lokapala_accountdb.t_ksatriya d ON c.ksatriya_id = d.ksatriya_id) as x GROUP BY ksatriya_id ORDER BY player_owned DESC")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var stat model.Most_ksa_owned
		err := result.Scan(&stat.Ksatriya_id, &stat.Ksatriya_name, &stat.Player_owned)
		if err != nil {
			panic(err.Error())
		}

		stats = append(stats, stat)

	}

	json.NewEncoder(w).Encode(stats)

}

func GetKsaTotalKda(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()

	var stats []model.Ksa_kda_stats

	game_mode := r.URL.Query().Get("game_mode")

	result, err := db.Query("SELECT a.ksatriya_id, b.ksatriya_name, SUM(`kill`) AS kill_count, SUM(death) AS death_count, SUM(assist) AS assist_count, ROUND(SUM(`kill`)/SUM(death),2) AS kill_death_rate FROM lokapala_roomdb.t_room_result_slot a JOIN lokapala_accountdb.t_ksatriya b ON a.ksatriya_id = b.ksatriya_id JOIN lokapala_roomdb.t_past_room r ON a.room_id = r.room_id JOIN lokapala_roomdb.t_room_result rr ON a.room_id = rr.room_id WHERE COALESCE(rr.match_id, 0) > 0 AND r.start_time >= '2020-05-19 17:00:00' AND IF(? = 0, TRUE, rr.game_mode = ?) AND NOT a.ksatriya_id = 901 GROUP BY b.ksatriya_id ORDER BY b.ksatriya_id ASC", game_mode, game_mode)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var stat model.Ksa_kda_stats
		err := result.Scan(&stat.Ksatriya_id, &stat.Ksatriya_name, &stat.Kill_count, &stat.Death_count, &stat.Assist_count, &stat.Kill_death_rate)
		if err != nil {
			panic(err.Error())
		}

		stats = append(stats, stat)

	}

	json.NewEncoder(w).Encode(stats)
}
