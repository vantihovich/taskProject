package handle

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"
	gr "github.com/vantihovich/taskProject/api"
	pg "github.com/vantihovich/taskProject/postgres"
	ps "github.com/vantihovich/taskProject/proto"
)

type reqUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginHandler struct {
	usp pg.UsersProvider
}

func NewLoginHandler(up *pg.UsersProvider) *LoginHandler {
	return &LoginHandler{usp: *up}
}

func (h *LoginHandler) Login(r *http.Request, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")

	params := reqUser{}

	err := json.NewDecoder(r.Body).Decode(&params)

	//TODO check if conditions are valid here
	if err == io.EOF {
		log.WithFields(log.Fields{"Error": err}).Info("Error empty request body")
		fmt.Fprint(w, "Please send a request body") //TODO handle the response
		return
	} else if err != nil {
		log.WithFields(log.Fields{"Error": err}).Info("Error occurred")
		http.Error(w, err.Error(), 500)
		return
	} else {
		//filling in "User" struct in postgres package and saving it to variable
		respusr, _ := h.usp.FindUserByEmailAndPassword(params.Email, params.Password) //will check later if err handling is valid here

		//Calling gRPS procedure with params from respusr variable
		resp, _ := gr.Cli.GenerateToken(context.Background(), //will check later if err handling is valid here
			&ps.Request{
				//Id:		  respusr.Id, --the parameter will be added after adding it to proto files
				Email:    respusr.Email,
				Password: respusr.Password,
			})
		log.WithFields(log.Fields{"Token": resp.Token, "Expires_at": resp.ExpiresAt}).Info("The server responded with:")
	}
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This string should be seen in browser")
}
