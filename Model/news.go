package model

type News struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	Release_date string `json:"release_date"`
	Type         int    `json:"type"`
}

type News_detail struct {
	News_id          int64  `json:"news_id"`
	Lang             string `json:"lang"`
	Title            string `json:"title"`
	Banner           string `json:"banner"`
	Banner_checksum  string `json:"banner_checksum"`
	Content          string `json:"content"`
	Content_checksum string `json:"content_checksum"`
}

type News_type struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type News_images struct {
	Id             int64  `json:"id"`
	Image_name     string `json:"image_name"`
	Image_checksum string `json:"image_checksum"`
	Uploader       string `json:"uploader"`
}

type News_img_fav struct {
	UserId  int64 `json:"user_id"`
	ImageId int64 `json:"image_id"`
}
