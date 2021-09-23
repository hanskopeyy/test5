package gacha

import (
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	controller "github.com/sen329/test5/Controller"
	model "github.com/sen329/test5/Model"
)

func AddGachaLoot(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO lokapala_accountdb.t_gacha_loot_table(gacha_id, gacha_item_id, chance, min_value, max_value) VALUES (?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	gacha_id := r.Form.Get("gacha_id")
	gacha_item_id := r.Form.Get("gacha_item_id")
	chance := r.Form.Get("chance")
	min_value := r.Form.Get("min_value")
	max_value := r.Form.Get("max_value")

	_, err = stmt.Exec(gacha_id, gacha_item_id, chance, min_value, max_value)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func GetAllGachaLoot(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	var gacha_loots []model.Gacha_loot_table

	result, err := db.Query("SELECT * from lokapala_accountdb.t_gacha_loot_table")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var gacha_loot model.Gacha_loot_table
		err := result.Scan(&gacha_loot.Gacha_id, &gacha_loot.Gacha_item_id, &gacha_loot.Chance, &gacha_loot.Min_value, &gacha_loot.Max_value)
		if err != nil {
			panic(err.Error())
		}

		gacha_loots = append(gacha_loots, gacha_loot)

	}

	json.NewEncoder(w).Encode(gacha_loots)

}

func GetGachaLoot(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id1 := r.URL.Query().Get("id1")
	id2 := r.URL.Query().Get("id2")

	var gacha_loot model.Gacha_loot_table
	results, err := db.Prepare("SELECT * from lokapala_accountdb.t_gacha_loot_table where gacha_id = ? AND gacha_item_id =? ")
	if err != nil {
		panic(err.Error())
	}

	result, err := results.Query(id1, id2)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {

		err := result.Scan(&gacha_loot.Gacha_id, &gacha_loot.Gacha_item_id, &gacha_loot.Chance, &gacha_loot.Min_value, &gacha_loot.Max_value)
		if err != nil {
			panic(err.Error())
		}

	}

	json.NewEncoder(w).Encode(gacha_loot)

}

func UpdateGachaLoot(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id1 := r.URL.Query().Get("id1")
	id2 := r.URL.Query().Get("id2")

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE lokapala_accountdb.t_gacha_loot_table SET gacha_id = ?, gacha_item_id = ?, chance = ?, min_value =?, max_value = ? where gacha_id = ? AND gacha_item_id = ?")
	if err != nil {
		panic(err.Error())
	}

	gacha_id := r.Form.Get("gacha_id")
	gacha_item_id := r.Form.Get("gacha_item_id")
	chance := r.Form.Get("chance")
	min_value := r.Form.Get("min_value")
	max_value := r.Form.Get("max_value")

	_, err = stmt.Exec(gacha_id, gacha_item_id, chance, min_value, max_value, id1, id2)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}

func DeleteGachaLoot(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id1 := r.URL.Query().Get("id1")
	id2 := r.URL.Query().Get("id2")

	stmt, err := db.Prepare("DELETE FROM lokapala_accountdb.t_gacha_loot_table WHERE gacha_id = ? AND gacha_item_id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(id1, id2)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}
