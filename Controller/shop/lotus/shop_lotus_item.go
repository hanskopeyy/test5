package lotus

import (
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	controller "github.com/sen329/test5/Controller"
	model "github.com/sen329/test5/Model"
)

func LotusAddNewItem(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO lokapala_accountdb.t_shop_lotus_item(item_type, item_id, amount, price, default_limit) VALUES (?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	item_type := r.Form.Get("item_type")
	item_id := r.Form.Get("item_id")
	amount := r.Form.Get("amount")
	price := r.Form.Get("price")
	default_limit := r.Form.Get("default_limit")

	_, err = stmt.Exec(item_type, item_id, amount, price, default_limit)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func LotusGetShopItems(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	var shop_items []model.Shop_lotus_item

	result, err := db.Query("SELECT * from lokapala_accountdb.t_shop_lotus_item")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var shop_item model.Shop_lotus_item
		err := result.Scan(&shop_item.Shop_lotus_item_id, &shop_item.Item_type, &shop_item.Item_id, &shop_item.Amount, &shop_item.Price, &shop_item.Default_limit)
		if err != nil {
			panic(err.Error())
		}

		shop_items = append(shop_items, shop_item)

	}

	json.NewEncoder(w).Encode(shop_items)

}

func LotusGetShopItem(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("id")

	var shop_item model.Shop_lotus_item
	result, err := db.Query("SELECT * from lokapala_accountdb.t_shop_lotus_item where shop_lotus_item_id = ?", id)
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

func LotusUpdateShopItem(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("id")

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE lokapala_accountdb.t_shop_lotus_item SET item_type = ?, item_id = ?, amount = ?, price = ?, default_limit = ? where shop_lotus_item_id = ?")
	if err != nil {
		panic(err.Error())
	}

	item_type := r.Form.Get("item_type")
	item_id := r.Form.Get("item_id")
	amount := r.Form.Get("amount")
	price := r.Form.Get("price")
	default_limit := r.Form.Get("default_limit")

	_, err = stmt.Exec(item_type, item_id, amount, price, default_limit, id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}

func LotusDeleteShopItem(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("id")

	stmt, err := db.Prepare("DELETE FROM lokapala_accountdb.t_shop_lotus_item WHERE shop_lotus_item_id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}
