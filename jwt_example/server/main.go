package main

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
)

var mySigningKey = []byte("mysupersecret")

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Super Secret Information")
}

func isAuthorized(endpoint func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] == nil {
			fmt.Fprintf(w, "Not Authorized")
		}

		token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("there was an error")
			}

			return mySigningKey, nil
		})
		if err != nil {
			fmt.Fprintf(w, err.Error())
		}

		fmt.Println(token.Claims)

		if token.Valid {
			endpoint(w, r)
		}
	})
}

func handleRequest() {
	http.Handle("/", isAuthorized(homePage))
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func main() {
	fmt.Println("My simple server")
	handleRequest()
}
