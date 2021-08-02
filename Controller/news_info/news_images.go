package newsinfo

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"

	controller "github.com/sen329/test5/Controller"
	model "github.com/sen329/test5/Model"
)

func AddImage(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO t_news_images(image_name,image_checksum,uploader) VALUES (?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	image_name, image_checksum, err := UploadFile(r, "uploadImage", "storage")
	if err != nil {
		panic(err)
	}

	uploader := r.Context().Value("user_id").(string)

	_, err = stmt.Exec(image_name, image_checksum, uploader)
	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode("Success")

}

func GetImages(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	var images []model.News_images

	result, err := db.Query("SELECT * FROM t_news_images")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var image model.News_images
		err := result.Scan(&image.Id, &image.Image_name, &image.Image_checksum, &image.Uploader)
		if err != nil {
			panic(err.Error())
		}

		images = append(images, image)
	}

	json.NewEncoder(w).Encode(images)
}

func GetImage(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("id")
	var image model.News_images

	result, err := db.Query("SELECT * FROM t_news_images WHERE id = ?", id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		err := result.Scan(&image.Id, &image.Image_name, &image.Image_checksum, &image.Uploader)
		if err != nil {
			panic(err.Error())
		}
	}

	json.NewEncoder(w).Encode(image)
}

func GetyourFavImages(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.Context().Value("user_id").(string)
	var images []model.News_images

	result, err := db.Query("SELECT img.* FROM t_news_images img INNER JOIN t_news_images_favorite img_fav on img.id = img_fav.imageid WHERE img_fav.userid = ?", id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var image model.News_images
		err := result.Scan(&image.Id, &image.Image_name, &image.Image_checksum, &image.Uploader)
		if err != nil {
			panic(err.Error())
		}

		images = append(images, image)
	}

	json.NewEncoder(w).Encode(images)
}

func DeleteImage(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("id")

	stmt, err := db.Prepare("DELETE FROM t_news_images WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}

	result, err := db.Query("SELECT * FROM t_news_images WHERE id = ?", id)
	if err != nil {
		panic(err.Error())
	}

	var image model.News_images

	for result.Next() {
		err := result.Scan(&image.Id, &image.Image_name, &image.Image_checksum, &image.Uploader)
		if err != nil {
			panic(err.Error())
		}
	}

	workdir, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}
	fileRandom := filepath.Join(workdir, image.Image_name)

	_, err = stmt.Exec(id)
	if err != nil {
		panic(err.Error())
	}
	os.Remove(fileRandom)

	json.NewEncoder(w).Encode("Success")
}
