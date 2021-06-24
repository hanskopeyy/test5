package shop

import (
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	model "github.com/sen329/test5/Model"
)

func AddShopBundle(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO t_shop_bundle(item_type, item_id, amount) VALUES (?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	item_type := r.Form.Get("item_type")
	item_id := r.Form.Get("item_id")
	amount := r.Form.Get("amount")

	_, err = stmt.Exec(item_type, item_id, amount)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func GetShopBundles(w http.ResponseWriter, r *http.Request) {
	var shop_bundles []model.Shop_bundle

	result, err := db.Query("SELECT * from t_shop_bundle")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var shop_bundle model.Shop_bundle
		err := result.Scan(&shop_bundle.Shop_id, &shop_bundle.Item_type, &shop_bundle.Item_id, &shop_bundle.Amount)
		if err != nil {
			panic(err.Error())
		}

		shop_bundles = append(shop_bundles, shop_bundle)

	}

	json.NewEncoder(w).Encode(shop_bundles)

}

func GetShopBundle(w http.ResponseWriter, r *http.Request) {
	id1 := r.URL.Query().Get("shop_id")
	id2 := r.URL.Query().Get("item_id")
	id3 := r.URL.Query().Get("item_type")

	var shop model.Shop_bundle
	results, err := db.Prepare("SELECT * from t_shop_bundle where shop_id = ? AND item_id = ? AND item_type = ?")
	if err != nil {
		panic(err.Error())
	}

	result, err := results.Query(id1, id2, id3)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {

		err := result.Scan(&shop.Shop_id, &shop.Item_id, &shop.Item_type, &shop.Amount)
		if err != nil {
			panic(err.Error())
		}

	}

	json.NewEncoder(w).Encode(shop)

}

func UpdateShopBundle(w http.ResponseWriter, r *http.Request) {
	id1 := r.URL.Query().Get("shop_id")
	id2 := r.URL.Query().Get("item_id")
	id3 := r.URL.Query().Get("item_type")

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE t_shop_bundle SET item_id = ?, item_type = ?, amount = ? where shop_id = ? AND item_id = ? AND item_type = ?")
	if err != nil {
		panic(err.Error())
	}

	item_type := r.Form.Get("item_type")
	item_id := r.Form.Get("item_id")
	amount := r.Form.Get("amount")

	_, err = stmt.Exec(item_id, item_type, amount, id1, id2, id3)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}

func DeleteShopBundle(w http.ResponseWriter, r *http.Request) {
	id1 := r.URL.Query().Get("shop_id")
	id2 := r.URL.Query().Get("item_id")
	id3 := r.URL.Query().Get("item_type")

	stmt, err := db.Prepare("DELETE FROM t_shop_bundle WHERE shop_id = ? AND item_id = ? AND item_type = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(id1, id2, id3)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}
