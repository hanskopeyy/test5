package newsinfo

import (
	"encoding/json"
	"net/http"

	controller "github.com/sen329/test5/Controller"
	model "github.com/sen329/test5/Model"
)

func AddtoFavorites(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()

	image_id := r.URL.Query().Get("id")
	user_id := r.Context().Value("user_id").(string)

	stmt, err := db.Prepare("INSERT INTO t_news_images_favorite(userId,imageId) VALUES (?,?)")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(user_id, image_id)
	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode("Success")
}

func GetAllFavorites(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	var images []model.News_img_fav

	result, err := db.Query("SELECT * FROM t_news_images_favorite")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var image model.News_img_fav
		err := result.Scan(&image.UserId, &image.ImageId)
		if err != nil {
			panic(err.Error())
		}

		images = append(images, image)
	}

	json.NewEncoder(w).Encode(images)

}

func GetYourFavorites(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	var images []model.News_img_fav

	user_id := r.Context().Value("user_id").(string)

	result, err := db.Query("SELECT * FROM t_news_images_favorite WHERE userid = ?", user_id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var image model.News_img_fav
		err := result.Scan(&image.UserId, &image.ImageId)
		if err != nil {
			panic(err.Error())
		}

		images = append(images, image)
	}

	json.NewEncoder(w).Encode(images)

}

func GetFavoritesById(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	var images []model.News_img_fav

	user_id := r.URL.Query().Get("id")

	result, err := db.Query("SELECT * FROM t_news_images_favorite WHERE userid = ?", user_id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var image model.News_img_fav
		err := result.Scan(&image.UserId, &image.ImageId)
		if err != nil {
			panic(err.Error())
		}

		images = append(images, image)
	}

	json.NewEncoder(w).Encode(images)

}

func RemoveFromFavorites(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	image_id := r.URL.Query().Get("id")
	user_id := r.Context().Value("user_id").(string)

	stmt, err := db.Prepare("DELETE FROM t_news_images_favorite WHERE userid = ? AND imageid = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(user_id, image_id)
	if err != nil {
		panic(err.Error())
	}

}
