package controller

//enter user stuff here

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"

	_ "github.com/go-sql-driver/mysql"
	model "github.com/sen329/test5/Model"

	"github.com/dgrijalva/jwt-go"
)

var db *sql.DB
var err error

var jwtKey = []byte("my_secret_key")

func Open() {
	db, err = sql.Open("mysql", "root:@/go_login_test")
	if err != nil {
		panic(err.Error())
	}

}

func Login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	email := r.Form.Get("email")
	password := r.Form.Get("password")

	stmt, err := db.Query("SELECT email, password FROM users WHERE email LIKE ?  ", email) //include role later
	if err != nil {
		panic(err.Error())
	}

	defer stmt.Close()

	var user model.User

	for stmt.Next() {
		err := stmt.Scan(&user.Email, &user.Password)
		if err != nil {
			panic(err.Error())
		}
	}

	check := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if check != nil {
		panic(check.Error())
	} else {
		fmt.Println("SUCCESSSSSSSSSSSSSSSSSSSSS")

		// fmt.Fprintf(w, "JWT Should be here as JSON")

		//Attempt #1

		// Declare the expiration time of the token
		// here, we have kept it as 5 minutes
		expirationTime := time.Now().Add(5 * time.Minute)
		// Create the JWT claims, which includes the username and expiry time

		claims := model.Claims{
			Email: user.Email,
			StandardClaims: jwt.StandardClaims{
				// In JWT, the expiry time is expressed as unix milliseconds
				ExpiresAt: expirationTime.Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			// If there is an error in creating the JWT return an internal server error
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		jwtToken := model.Token{
			Token: tokenString,
		}

		json.NewEncoder(w).Encode(jwtToken)

		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
		})

	}

	defer db.Close()

}

func Register(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}
	stmt, err := db.Prepare("INSERT INTO USERS(name, email, password) VALUES (?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	name := r.Form.Get("name")
	email := r.Form.Get("email")
	password := r.Form.Get("password")

	var pwd = []byte(password)

	pwdhash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(name, email, pwdhash)
	if err != nil {
		panic(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")

	defer db.Close()
}

func Refresh(w http.ResponseWriter, r *http.Request) {
	// (BEGIN) The code uptil this point is the same as the first part of the `Welcome` route
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tknStr := c.Value
	claims := model.Claims{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	// (END) The code up-till this point is the same as the first part of the `Welcome` route

	// We ensure that a new token is not issued until enough time has elapsed
	// In this case, a new token will only be issued if the old token is within
	// 30 seconds of expiry. Otherwise, return a bad request status
	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Now, create a new token for the current use, with a renewed expiration time
	expirationTime := time.Now().Add(5 * time.Minute)
	claims.ExpiresAt = expirationTime.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Set the new token as the users `token` cookie
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	defer db.Close()
}