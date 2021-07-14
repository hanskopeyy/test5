package controller

import (
	"encoding/json"
	"net/http"
)

func ShopAddItem(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkshop(user_id, role_id) {
		AddShopItem(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}

func ShopGetAllItems(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkshop(user_id, role_id) {
		GetShopItems(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}

func ShopGetItem(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkshop(user_id, role_id) {
		GetShopItem(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}

func ShopUpdateItem(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkshop(user_id, role_id) {
		UpdateShopItem(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}

func ShopDeleteItem(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkshop(user_id, role_id) {
		DeleteShopItem(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}

func ShopAddBundle(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkshop(user_id, role_id) {
		AddShopBundle(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}

func ShopGetAllBundles(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkshop(user_id, role_id) {
		GetShopBundles(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}

func ShopGetBundle(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkshop(user_id, role_id) {
		GetShopBundle(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}

func ShopUpdateBundle(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkshop(user_id, role_id) {
		UpdateShopBundle(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}

func ShopDeleteBundle(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkshop(user_id, role_id) {
		DeleteShopBundle(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}
