package customware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

type Key string

func Authenticate() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				tokenStrings := r.Header.Values("jwt-auth-key")
				fmt.Println("path is", r.URL.Path)
				if tokenStrings == nil {
					// ctx := context.WithValue(r.Context(), "next_api", next)

					ctx := context.WithValue(r.Context(), "after_login", r.URL.Path)
					http.Redirect(w, r.WithContext(ctx), "/login", http.StatusSeeOther) // should we redirect or should we make a router that make login here??
					//error with Redirect: Error: Exceeded maxRedirects. Probably stuck in a redirect loop http://127.0.0.1:3000/login
					//because i put auth middleware before login route so it created an infinity loop
					return
				}
				fmt.Print(tokenStrings)
				tokenString := tokenStrings[0]
				fmt.Print(viper.GetString("secret"))
				token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) { return []byte(viper.GetString("secret")), nil })
				// also func here is to return a key for comparing, which is our secret
				// we need config file to store secret!
				if err != nil {
					w.WriteHeader(500)
					w.Write([]byte(err.Error()))
				}
				userInfo := token.Claims
				if valid := token.Valid; valid {
					ctx := context.WithValue(r.Context(), Key("user"), userInfo)
					next.ServeHTTP(w, r.WithContext(ctx))
				} else {
					w.Write([]byte("invalid token,try again"))
					return
				}
			})
	}
}
