package controller

import (
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	model "github.com/sen329/test5/Model"
)

func GenerateVoucher(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("call lokapala_melonpaymentdb.p_generate_key(?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	count := r.Form.Get("count")
	voucher_id := r.Form.Get("voucher_id")
	expire_date := r.Form.Get("expire_date")

	_, err = stmt.Exec(count, voucher_id, expire_date)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func GetAllVouchers(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	var vouchers []model.Voucher
	result, err := db.Query("SELECT * FROM lokapala_melonpaymentdb.t_voucher")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var voucher model.Voucher
		err := result.Scan(&voucher.Id, &voucher.Key, &voucher.Created_date, &voucher.User_id)
		if err != nil {
			panic(err.Error())
		}
		vouchers = append(vouchers, voucher)
	}

	json.NewEncoder(w).Encode(vouchers)
}

func GetVoucher(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	id := r.URL.Query().Get("id")

	var voucher model.Voucher
	result, err := db.Query("SELECT * FROM lokapala_melonpaymentdb.t_voucher WHERE id = ?", id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		err := result.Scan(&voucher.Id, &voucher.Key, &voucher.Created_date, &voucher.User_id)
		if err != nil {
			panic(err.Error())
		}
	}

	json.NewEncoder(w).Encode(voucher)
}

func UpdateVoucher(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	id := r.URL.Query().Get("id")

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE lokapala_melonpaymentdb.t_voucher SET voucher_id = ? WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}

	voucher_id := r.Form.Get("voucher_id")

	_, err = stmt.Exec(voucher_id, id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func DeleteVoucher(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	id := r.URL.Query().Get("id")

	stmt, err := db.Prepare("DELETE FROM lokapala_melonpaymentdb.t_voucher WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func AddVoucher(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO lokapala_melonpaymentdb.t_voucher_detail(voucher_id, item_type, item_id, amount, detail) VALUES (?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	voucher_id := r.Form.Get("voucher_id")
	item_id := r.Form.Get("item_id")
	item_type := r.Form.Get("item_type")
	amount := r.Form.Get("amount")
	detail := r.Form.Get("detail")

	_, err = stmt.Exec(voucher_id, item_id, item_type, amount, detail)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func GetAllVoucherDetails(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	var voucher_details []model.Voucher_detail
	result, err := db.Query("SELECT * FROM lokapala_melonpaymentdb.t_voucher_detail")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var voucher_detail model.Voucher_detail
		err := result.Scan(&voucher_detail.Voucher_id, &voucher_detail.Item_type, &voucher_detail.Item_id, &voucher_detail.Amount, &voucher_detail.Detail)
		if err != nil {
			panic(err.Error())
		}
		voucher_details = append(voucher_details, voucher_detail)
	}

	json.NewEncoder(w).Encode(voucher_details)
}

func GetVoucherDetail(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	voucher_id := r.URL.Query().Get("voucher_id")

	var voucher_detail model.Voucher_detail
	result, err := db.Query("SELECT * FROM lokapala_melonpaymentdb.t_voucher_detail WHERE voucher_id = ?", voucher_id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		err := result.Scan(&voucher_detail.Voucher_id, &voucher_detail.Item_type, &voucher_detail.Item_id, &voucher_detail.Amount, &voucher_detail.Detail)
		if err != nil {
			panic(err.Error())
		}
	}

	json.NewEncoder(w).Encode(voucher_detail)
}

func UpdateVoucherDetail(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	voucher_id := r.URL.Query().Get("voucher_id")

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE lokapala_melonpaymentdb.t_voucher_detail SET item_id = ?, item_type = ?, amount = ? WHERE voucher_id = ?")
	if err != nil {
		panic(err.Error())
	}

	item_id := r.Form.Get("item_id")
	item_type := r.Form.Get("item_type")
	amount := r.Form.Get("amount")

	_, err = stmt.Exec(item_id, item_type, amount, voucher_id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func DeleteVoucherDetail(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	voucher_id := r.URL.Query().Get("voucher_id")

	stmt, err := db.Prepare("DELETE FROM lokapala_melonpaymentdb.t_voucher_detail WHERE voucher_id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(voucher_id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}
