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

var JwtKey = []byte(goDotEnvVariable("SECRET_KEY"))

func Login(w http.ResponseWriter, r *http.Request) {
	db := OpenGMAdmin()
	defer db.Close()
	w.Header().Add("Content-Type", "application/json")
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	email := r.Form.Get("email")
	password := r.Form.Get("password")

	stmt, err := db.Query("SELECT users.id, users.email, users.password, roles.id FROM users LEFT JOIN users_roles ON users.id = users_roles.user_id LEFT JOIN roles ON users_roles.role_id = roles.id WHERE email LIKE ?  ", email)
	if err != nil {
		panic(err.Error())
	}

	defer stmt.Close()

	var user model.User
	var role model.Roles

	for stmt.Next() {
		err := stmt.Scan(&user.Id, &user.Email, &user.Password, &role.Id)
		if err != nil {
			panic(err.Error())
		}
	}

	check := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if check != nil {
		json.NewEncoder(w).Encode("Wrong Password")
		panic(check.Error())
	} else {
		// fmt.Println("SUCCESSSSSSSSSSSSSSSSSSSSS")

		// fmt.Fprintf(w, "JWT Should be here as JSON")

		//Attempt #1

		// Declare the expiration time of the token
		// here, we have kept it as 5 minutes
		expirationTime := time.Now().Add(60 * time.Minute)
		// Create the JWT claims, which includes the username and expiry time

		claims := model.Claims{
			Email:   user.Email,
			User_id: user.Id,
			Role_id: role.Id,
			StandardClaims: jwt.StandardClaims{
				// In JWT, the expiry time is expressed as unix milliseconds
				ExpiresAt: expirationTime.Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		tokenString, err := token.SignedString(JwtKey)
		if err != nil {
			// If there is an error in creating the JWT return an internal server error
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		jwtToken := model.Token{
			Token: tokenString,
		}

		json.NewEncoder(w).Encode(jwtToken)

		// http.SetCookie(w, &http.Cookie{
		// 	Name:    "token",
		// 	Value:   tokenString,
		// 	Expires: expirationTime,
		// })

	}

}

func Register(w http.ResponseWriter, r *http.Request) {
	db := OpenGMAdmin()
	defer db.Close()
	w.Header().Add("Content-Type", "application/json")
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}
	stmt, err := db.Prepare("INSERT INTO users(name, email, password) VALUES (?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	name := r.Form.Get("name")
	email := r.Form.Get("email")
	password := r.Form.Get("password")
	role_id := r.Form.Get("role_id")

	var pwd = []byte(password)

	pwdhash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(name, email, pwdhash)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(role_id)

	if len(role_id) > 0 {
		stmt2, err := db.Prepare("INSERT INTO users_roles(user_id, role_id) VALUES (?,?)")
		if err != nil {
			panic(err.Error())
		}

		stmt3, err := db.Query("SELECT id FROM users WHERE email LIKE ?  ", email)
		if err != nil {
			panic(err.Error())
		}

		var user model.User

		for stmt3.Next() {
			err := stmt3.Scan(&user.Id)
			if err != nil {
				panic(err.Error())
			}
		}

		_, err = stmt2.Exec(user.Id, role_id)
		if err != nil {
			panic(err.Error())
		}
	}

	w.Header().Set("Content-Type", "application/json")

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
		return JwtKey, nil
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
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Set the new token as the users `token` cookie
	// http.SetCookie(w, &http.Cookie{
	// 	Name:    "token",
	// 	Value:   tokenString,
	// 	Expires: expirationTime,
	// })

	jwtToken := model.Token{
		Token: tokenString,
	}

	json.NewEncoder(w).Encode(jwtToken)

	defer db.Close()
}

// func Checktest(w http.ResponseWriter, r *http.Request) {
// 	user_id := r.Context().Value("user_id").(string)
// 	role_id := r.Context().Value("role_id").(string)
// 	if Checkuser(user_id, role_id) {
// 		Test(w, r)
// 		//here the main code for anything
// 	} else {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		json.NewEncoder(w).Encode("Not authorized")
// 	}

// 	// fmt.Fprintf(w, r.Context().Value("user_id").(string))
// 	// fmt.Fprintf(w, r.Context().Value("role_id").(string))
// }

func Test(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Authorize works")
}

func NewRole(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Authorize works")
}
