package controller

import (
	"github.com/secsy/goftp"
)

func FTP() (connect *goftp.Client) {
	address := goDotEnvVariable("FTP_ADDR")
	config := goftp.Config{
		User:     goDotEnvVariable("FTP_USERNAME"),
		Password: goDotEnvVariable("FTP_PASSWORD"),
	}

	connect, err := goftp.DialConfig(config, address)
	if err != nil {
		panic(err.Error())
	}
	return connect
}
