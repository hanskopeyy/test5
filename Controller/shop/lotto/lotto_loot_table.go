package lotto

import (
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	controller "github.com/sen329/test5/Controller"
	model "github.com/sen329/test5/Model"
)

func AddlottoLoot(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO lokapala_accountdb.t_lotto_loot_table(lotto_id, lotto_item_id, amount) VALUES (?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	lotto_id := r.Form.Get("lotto_id")
	lotto_item_id := r.Form.Get("lotto_item_id")
	amount := r.Form.Get("amount")

	_, err = stmt.Exec(lotto_id, lotto_item_id, amount)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func GetlottoLoots(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	var l_loots []model.Lotto_loot_table
	result, err := db.Query("SELECT * FROM lokapala_accountdb.t_lotto_loot_table")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var l_loot model.Lotto_loot_table
		err := result.Scan(&l_loot.Lotto_id, &l_loot.Lotto_item_id, &l_loot.Amount)
		if err != nil {
			panic(err.Error())
		}

		l_loots = append(l_loots, l_loot)
	}

	json.NewEncoder(w).Encode(l_loots)
}

func GetlottoLoot(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	l_id := r.URL.Query().Get("id")
	l_item_id := r.URL.Query().Get("item_id")

	var l_loot model.Lotto_loot_table
	results, err := db.Prepare("SELECT * from lokapala_accountdb.t_lotto_loot_table WHERE lotto_id = ? AND lotto_item_id = ?")
	if err != nil {
		panic(err.Error())
	}

	result, err := results.Query(l_id, l_item_id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		err := result.Scan(&l_loot.Lotto_id, &l_loot.Lotto_item_id, &l_loot.Amount)
		if err != nil {
			panic(err.Error())
		}
	}

	json.NewEncoder(w).Encode(l_loot)
}

func GetlottoLootByLottoId(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("id")

	var l_loot model.Lotto_loot_table
	result, err := db.Query("SELECT * FROM lokapala_accountdb.t_lotto_loot_table WHERE lotto_id = ?", id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		err := result.Scan(&l_loot.Lotto_id, &l_loot.Lotto_item_id, &l_loot.Amount)
		if err != nil {
			panic(err.Error())
		}
	}
	json.NewEncoder(w).Encode(l_loot)
}

func UpdatelottoLoot(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	l_id := r.URL.Query().Get("id")
	l_item_id := r.URL.Query().Get("item_id")

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE lokapala_accountdb.t_lotto_loot_table SET amount = ? WHERE lotto_id = ? AND lotto_item_id = ?")
	if err != nil {
		panic(err.Error())
	}

	amount := r.Form.Get("amount")

	_, err = stmt.Exec(amount, l_id, l_item_id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func DeletelottoLoot(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	l_id := r.URL.Query().Get("id")
	l_item_id := r.URL.Query().Get("item_id")

	stmt, err := db.Prepare("DELETE FROM lokapala_accountdb.t_lotto_loot_table WHERE lotto_id = ? AND lotto_item_id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(l_id, l_item_id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}
