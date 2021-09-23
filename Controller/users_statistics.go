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
