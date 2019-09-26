package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
)

func init() {
	//global.tokenAuth = jwtauth.New("HS256", []byte(global.key), nil)

}
func main() {
	r := chi.NewRouter()
	r.Group(func(r chi.Router) {

		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			//_, claims, _ := jwtauth.FromContext(r.Context())
			w.Write([]byte("Welcome to home page "))
		})
	})
	r.Group(func(r chi.Router) {

		r.Get("/login", func(w http.ResponseWriter, r *http.Request) {
			bodyBytes, err := ioutil.ReadAll(r.Body)
			if err != nil {
				log.Fatal(err)
			}
			var tp Credentials
			json.Unmarshal(bodyBytes, &tp)
			if _, ok := users[tp.Username]; ok {
				if users[tp.Username] == tp.Password {
					expirationTime := time.Now().Add(5 * time.Minute)

					http.SetCookie(w, &http.Cookie{
						Name:    "jwt",
						Value:   tokengen(tp.Username),
						Expires: expirationTime,
					})
					w.Write([]byte("Login successfull  -" + tp.Username + "\n token  is " + tokengen(tp.Username)))

				} else {
					w.Write([]byte("Invalid password  -" + tp.Username))
				}
			} else {
				w.Write([]byte("Not a registred User  -" + tp.Username))
			}

		})
	})
	r.Group(func(r chi.Router) {
		// Seek, verify and validate JWT tokens

		r.Use(jwtauth.Verifier(tokenAuth))
		// Handle valid / invalid tokens. In this example, we use
		// the provided authenticator middleware, but you can write your
		// own very easily, look at the Authenticator method in jwtauth.go
		// and tweak it, its not scary.
		r.Use(jwtauth.Authenticator)

		r.Get("/admin", func(w http.ResponseWriter, r *http.Request) {
			_, claims, _ := jwtauth.FromContext(r.Context())
			w.Write([]byte(fmt.Sprintf("protected area. hi %v", claims["user_id"])))
		})
	})
}
