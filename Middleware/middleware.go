package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"

	controller "github.com/sen329/test5/Controller"
	model "github.com/sen329/test5/Model"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization, X-Requested-With")
		w.Header().Set("Access-Control-Allow-Methods", "GET, HEAD, POST, PUT, OPTIONS, DELETE")
		w.Header().Set("Accept", "*/*")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			fmt.Println("Malformed token")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Malformed Token"))
		} else {
			jwtToken := authHeader[1]

			token, err := jwt.ParseWithClaims(jwtToken, &model.Claims{}, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(controller.JwtKey), nil
			})

			if claims, ok := token.Claims.(*model.Claims); ok && token.Valid {
				ctx := context.WithValue(context.Background(), "user_id", claims.User_id)
				ctx = context.WithValue(ctx, "role_id", claims.Role_id)
				next.ServeHTTP(w, r.WithContext(ctx))

			} else {
				fmt.Println(err)
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Unauthorized"))
			}
		}
	})
}

func CheckRoleShop(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user_id := r.Context().Value("user_id").(string)
		role_id := r.Context().Value("role_id").(string)
		if Checkshop(user_id, role_id) {
			next.ServeHTTP(w, r)
		} else {
			fmt.Println("Role doesn't match")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Not authorized")
		}
	})
}

func CheckRoleMail(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user_id := r.Context().Value("user_id").(string)
		role_id := r.Context().Value("role_id").(string)
		if Checkmail(user_id, role_id) {
			next.ServeHTTP(w, r)
		} else {
			fmt.Println("Role doesn't match")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Not authorized")
		}
	})
}

func CheckRoleUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user_id := r.Context().Value("user_id").(string)
		role_id := r.Context().Value("role_id").(string)
		if Checkuser(user_id, role_id) {
			next.ServeHTTP(w, r)
		} else {
			fmt.Println("Role doesn't match")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Not authorized")
		}
	})
}

func CheckRolePlayer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user_id := r.Context().Value("user_id").(string)
		role_id := r.Context().Value("role_id").(string)
		if Checkplayer(user_id, role_id) {
			next.ServeHTTP(w, r)
		} else {
			fmt.Println("Role doesn't match")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Not authorized")
		}
	})
}

func CheckRoleKsaRot(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user_id := r.Context().Value("user_id").(string)
		role_id := r.Context().Value("role_id").(string)
		if Check_ksa_rot(user_id, role_id) {
			next.ServeHTTP(w, r)
		} else {
			fmt.Println("Role doesn't match")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Not authorized")
		}
	})
}

func CheckRoleMatches(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user_id := r.Context().Value("user_id").(string)
		role_id := r.Context().Value("role_id").(string)
		if Check_matches(user_id, role_id) {
			next.ServeHTTP(w, r)
		} else {
			fmt.Println("Role doesn't match")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Not authorized")
		}
	})
}

func CheckRoleReport(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user_id := r.Context().Value("user_id").(string)
		role_id := r.Context().Value("role_id").(string)
		if Check_player_report(user_id, role_id) {
			next.ServeHTTP(w, r)
		} else {
			fmt.Println("Role doesn't match")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Not authorized")
		}
	})
}

func CheckRoleBlacklist(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user_id := r.Context().Value("user_id").(string)
		role_id := r.Context().Value("role_id").(string)
		if Check_blacklist(user_id, role_id) {
			next.ServeHTTP(w, r)
		} else {
			fmt.Println("Role doesn't match")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Not authorized")
		}
	})
}

func CheckRoleVoucher(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user_id := r.Context().Value("user_id").(string)
		role_id := r.Context().Value("role_id").(string)
		if Check_voucher(user_id, role_id) {
			next.ServeHTTP(w, r)
		} else {
			fmt.Println("Role doesn't match")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Not authorized")
		}
	})
}

func CheckRoleJudge(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user_id := r.Context().Value("user_id").(string)
		role_id := r.Context().Value("role_id").(string)
		if Check_judge(user_id, role_id) {
			next.ServeHTTP(w, r)
		} else {
			fmt.Println("Role doesn't match")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Not authorized")
		}
	})
}

func CheckRolePlayerStats(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user_id := r.Context().Value("user_id").(string)
		role_id := r.Context().Value("role_id").(string)
		if Check_player_stats(user_id, role_id) {
			next.ServeHTTP(w, r)
		} else {
			fmt.Println("Role doesn't match")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Not authorized")
		}
	})
}
