package lotus

import (
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	controller "github.com/sen329/test5/Controller"
	model "github.com/sen329/test5/Model"
)

func AddLotusPeriod(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO lokapala_accountdb.t_shop_lotus_period(start_date, end_date) VALUES (?,?)")
	if err != nil {
		panic(err.Error())
	}

	start_date := r.Form.Get("start_date")
	end_date := r.Form.Get("end_date")

	_, err = stmt.Exec(start_date, end_date)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func LotusGetShopPeriods(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	var shop_periods []model.Shop_lotus_period

	result, err := db.Query("SELECT * from lokapala_accountdb.t_shop_lotus_period")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var shop_period model.Shop_lotus_period
		err := result.Scan(&shop_period.Shop_lotus_period_id, &shop_period.Start_date, &shop_period.End_date)
		if err != nil {
			panic(err.Error())
		}

		shop_periods = append(shop_periods, shop_period)

	}

	json.NewEncoder(w).Encode(shop_periods)

}

func LotusGetShopPeriod(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("id")

	var shop_item model.Shop_lotus_item
	result, err := db.Query("SELECT * from lokapala_accountdb.t_shop_lotus_period where shop_lotus_period_id = ?", id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {

		err := result.Scan(&shop_item.Shop_lotus_item_id, &shop_item.Item_type, &shop_item.Item_id, &shop_item.Amount, &shop_item.Price, &shop_item.Default_limit)
		if err != nil {
			panic(err.Error())
		}

	}

	json.NewEncoder(w).Encode(shop_item)

}

func LotusUpdateShopPeriod(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("id")

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE lokapala_accountdb.t_shop_lotus_period SET start_date = ?, end_date = ? where shop_lotus_period_id = ?")
	if err != nil {
		panic(err.Error())
	}

	start_date := r.Form.Get("start_date")
	end_date := r.Form.Get("end_date")

	_, err = stmt.Exec(start_date, end_date, id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}

func LotusDeleteShopPeriod(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("id")

	stmt, err := db.Prepare("DELETE FROM lokapala_accountdb.t_shop_lotus_period WHERE shop_lotus_period_id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}
