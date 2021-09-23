package lotus

import (
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	controller "github.com/sen329/test5/Controller"
	model "github.com/sen329/test5/Model"
)

func AddLotus(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO lokapala_accountdb.t_shop_lotus(shop_lotus_period_id, shop_lotus_item_id, player_limit) VALUES (?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	shop_lotus_period_id := r.Form.Get("shop_lotus_period_id")
	shop_lotus_item_id := r.Form.Get("shop_lotus_item_id")
	player_limit := r.Form.Get("player_limit")

	_, err = stmt.Exec(shop_lotus_period_id, shop_lotus_item_id, player_limit)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func GetAllLotus(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	var shop_lotuss []model.Shop_lotus

	result, err := db.Query("SELECT * from lokapala_accountdb.t_shop_lotus")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var shop_lotus model.Shop_lotus
		err := result.Scan(&shop_lotus.Shop_lotus_period_id, &shop_lotus.Shop_lotus_item_id, &shop_lotus.Player_limit)
		if err != nil {
			panic(err.Error())
		}

		shop_lotuss = append(shop_lotuss, shop_lotus)

	}

	json.NewEncoder(w).Encode(shop_lotuss)

}

func GetLotus(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id1 := r.URL.Query().Get("id1")
	id2 := r.URL.Query().Get("id2")

	var shop_lotus model.Shop_lotus
	results, err := db.Prepare("SELECT * from lokapala_accountdb.t_shop_lotus where shop_lotus_period_id = ? AND shop_lotus_item_id =? ")
	if err != nil {
		panic(err.Error())
	}

	result, err := results.Query(id1, id2)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {

		err := result.Scan(&shop_lotus.Shop_lotus_period_id, &shop_lotus.Shop_lotus_item_id, &shop_lotus.Player_limit)
		if err != nil {
			panic(err.Error())
		}

	}

	json.NewEncoder(w).Encode(shop_lotus)

}

func UpdateLotusShop(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id1 := r.URL.Query().Get("id1")
	id2 := r.URL.Query().Get("id2")

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE lokapala_accountdb.t_shop_lotus SET shop_lotus_period_id = ?, shop_lotus_item_id = ?, player_limit = ? where shop_lotus_item_id = ? AND shop_lotus_period_id = ?")
	if err != nil {
		panic(err.Error())
	}

	shop_lotus_period_id := r.Form.Get("start_date")
	shop_lotus_item_id := r.Form.Get("end_date")
	player_limit := r.Form.Get("player_limit")

	_, err = stmt.Exec(shop_lotus_period_id, shop_lotus_item_id, player_limit, id1, id2)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}

func DeleteLotusShop(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id1 := r.URL.Query().Get("id1")
	id2 := r.URL.Query().Get("id2")

	stmt, err := db.Prepare("DELETE FROM lokapala_accountdb.t_shop_lotus WHERE shop_lotus_period_id = ? AND shop_lotus_item_id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(id1, id2)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}
