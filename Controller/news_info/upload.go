package newsinfo

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"io"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	controller "github.com/sen329/test5/Controller"
	model "github.com/sen329/test5/Model"
)

const latin = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz01233456789"

func UploadFile(r *http.Request, form string, paths ...string) (string, string, error) {
	err := r.ParseMultipartForm(10 * 1024 * 1024)
	if err != nil {
		panic(err)
	}

	uploadedFile, fileHandler, err := r.FormFile(form)
	if err != nil {
		panic(err.Error())
	}
	defer uploadedFile.Close()

	// Checksum
	hash := md5.New()
	if _, err := io.Copy(hash, uploadedFile); err != nil {
		panic(err.Error())
	}
	hashInBytes := hash.Sum(nil)[:16]
	checksum := hex.EncodeToString(hashInBytes)

	// Random Name . Extension
	buffer := randomName(7)
	fileExtension := strings.Split(fileHandler.Filename, ".")
	newRandName := buffer + "." + fileExtension[len(fileExtension)-1]

	// Get Path
	fileLocation, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}
	for _, path := range paths {
		fileLocation = filepath.Join(fileLocation, path)
	}
	fileLocation = filepath.Join(fileLocation, newRandName)

	locate, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err.Error())
	}

	io.Copy(locate, uploadedFile)

	return newRandName, checksum, nil
}

func randomName(digit int) string {
	db := controller.Open()
	defer db.Close()

	var buffer bytes.Buffer
	for i := 0; i < digit; i++ {
		buffer.WriteString(string(latin[rand.Intn(len(latin))]))
	}
	newRandName := string(buffer.String())

	// Check if exist
	result, err := db.Query("SELECT * FROM t_news_images WHERE image_name LIKE ?", newRandName+"%")
	if err != nil {
		panic(err.Error())
	}
	var images []model.News_images
	for result.Next() {
		var image model.News_images
		err := result.Scan(&image.Id, &image.Image_name, &image.Image_checksum, &image.Uploader)
		if err != nil {
			panic(err.Error())
		}

		images = append(images, image)
	}

	if len(images) != 0 {
		// fmt.Println(newRandName)
		// fmt.Println("Dupe found, re-randomizing name")
		newRandName = randomName(7)
	} else {
		// fmt.Println("Success no dupe")
	}
	return newRandName
}

func CheckorUpload(r *http.Request, form string) (string, string, error) {
	db := controller.Open()
	defer db.Close()
	err := r.ParseMultipartForm(10 * 1024 * 1024)
	if err != nil {
		panic(err)
	}

	getFromForm := r.Form.Get(form)
	var checksum string
	if len(getFromForm) == 0 {
		getFromForm, checksum, err := UploadFile(r, form, "storage")
		if err != nil {
			panic(err)
		}

		stmt, err := db.Prepare("INSERT INTO t_news_images(image_name,image_checksum,uploader) VALUES (?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		uploader := r.Context().Value("user_id").(string)

		_, err = stmt.Exec(getFromForm, checksum, uploader)
		if err != nil {
			panic(err)
		}
		return getFromForm, checksum, err
	}
	result, err := db.Query("SELECT * FROM t_news_images WHERE image_name = ?", getFromForm)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var image model.News_images
		err := result.Scan(&image.Id, &image.Image_name, &image.Image_checksum, &image.Uploader)
		if err != nil {
			panic(err.Error())
		}
		getFromForm = image.Image_name
		checksum = image.Image_checksum
	}
	return getFromForm, checksum, err
}
