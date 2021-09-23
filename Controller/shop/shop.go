package shop

import (
	"database/sql"
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	controller "github.com/sen329/test5/Controller"
	model "github.com/sen329/test5/Model"
)

var db *sql.DB

func AddShopItem(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO lokapala_accountdb.t_shop(item_id, item_type, amount, price_coin, price_citrine, price_lotus, release_date, description) VALUES (?,?,?,?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	item_id := r.Form.Get("item_id")
	item_type := r.Form.Get("item_type")
	amount := r.Form.Get("amount")
	price_coin := r.Form.Get("price_coin")
	price_citrine := r.Form.Get("price_citrine")
	price_lotus := r.Form.Get("price_lotus")
	release_date := r.Form.Get("release_date")
	var description string = r.Form.Get("description")

	_, err = stmt.Exec(item_id, item_type, amount, price_coin, price_citrine, price_lotus, release_date, description)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func GetShopItems(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()

	var shops []model.Shop
	result, err := db.Query("SELECT * from lokapala_accountdb.t_shop")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var shop model.Shop
		err := result.Scan(&shop.Shop_id, &shop.Item_id, &shop.Item_type, &shop.Amount, &shop.Price_coin, &shop.Price_citrine, &shop.Price_lotus, &shop.Release_date, &shop.Description)
		if err != nil {
			panic(err.Error())
		}

		shops = append(shops, shop)

	}

	json.NewEncoder(w).Encode(shops)

}

func GetShopItem(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("id")

	var shop model.Shop
	result, err := db.Query("SELECT * from lokapala_accountdb.t_shop where shop_id = ?", id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {

		err := result.Scan(&shop.Shop_id, &shop.Item_id, &shop.Item_type, &shop.Amount, &shop.Price_coin, &shop.Price_citrine, &shop.Price_lotus, &shop.Release_date, &shop.Description)
		if err != nil {
			panic(err.Error())
		}

	}

	json.NewEncoder(w).Encode(shop)

}

func UpdateShopItem(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("id")

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE lokapala_accountdb.t_shop SET item_id = ?, item_type = ?, amount = ?, price_coin = ?, price_citrine = ?, price_lotus = ?, release_date =?, description =?  where shop_id = ?")
	if err != nil {
		panic(err.Error())
	}

	item_id := r.Form.Get("item_id")
	item_type := r.Form.Get("item_type")
	amount := r.Form.Get("amount")
	price_coin := r.Form.Get("price_coin")
	price_citrine := r.Form.Get("price_citrine")
	price_lotus := r.Form.Get("price_lotus")
	release_date := r.Form.Get("release_date")
	description := r.Form.Get("description")

	_, err = stmt.Exec(item_id, item_type, amount, price_coin, price_citrine, price_lotus, release_date, description, id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}

func DeleteShopItem(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("id")

	stmt, err := db.Prepare("DELETE FROM lokapala_accountdb.t_shop WHERE shop_id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}
