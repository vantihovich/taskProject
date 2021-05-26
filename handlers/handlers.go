package handle

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"
	gr "github.com/vantihovich/taskProject/api"
	ps "github.com/vantihovich/taskProject/proto"
)

type user struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Hello(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "This string should be seen in browser")

}

func Login(l http.ResponseWriter, k *http.Request) {
	l.Header().Set("Content-Type", "application/json")

	params := user{}

	err := json.NewDecoder(k.Body).Decode(&params)

	if err == io.EOF {
		log.WithFields(log.Fields{"Error": err}).Info("Error empty request body")
		fmt.Fprint(l, "Please send a request body")
		return
	} else if err != nil {
		log.WithFields(log.Fields{"Error": err}).Info("Error occurred")
		http.Error(l, err.Error(), 500)
		return
	}

	fmt.Fprint(l, "The request to login with parameters:", params)
	log.WithFields(log.Fields{"Parameters": params}).Info("The request to login with parameters:")

	email := params.Email
	password := params.Password
	fmt.Println("the user:", email)
	fmt.Println("the password:", password)

	resp, err2 := gr.Cli.GenerateToken(context.Background(),
		&ps.Request{
			Email:    email,
			Password: password,
		})

	if err2 != nil {
		log.WithFields(log.Fields{"Error": err2}).Panicf("could not get answer:")
	}
	log.WithFields(log.Fields{"Token": resp.Token, "Expires_at": resp.ExpiresAt}).Info("could not get answer:")
	fmt.Fprint(l, "Token and expires_at are:", resp.Token, resp.ExpiresAt)
}
