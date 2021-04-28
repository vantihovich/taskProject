package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	gr "github.com/vantihovich/taskProject/api"
	ps "github.com/vantihovich/taskProject/proto"
)

var cli ps.GetCredsClient

func login(l http.ResponseWriter, k *http.Request) {
	l.Header().Set("Content-Type", "application/json")

	type User struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	params := User{}

	err := json.NewDecoder(k.Body).Decode(&params)

	if err == io.EOF {
		log.Println("The error:", err)
		fmt.Fprint(l, "Please send a request body")
		return
	} else if err != nil {
		log.Println("The error", err)
		http.Error(l, err.Error(), 500)
		return
	}

	fmt.Fprint(l, "запрос на логин с параметрами:", params)

	email := params.Email
	password := params.Password
	fmt.Println("the user:", email)
	fmt.Println("the password:", password)

	// fmt.Println("Старт gRPC клиента")

	// conn, err := grpc.Dial("127.0.0.1:3500", grpc.WithInsecure())
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// client := ps.NewGetCredsClient(conn)

	resp, err2 := cli.GenerateToken(context.Background(),
		&ps.Request{
			Email:    email,
			Password: password,
		})

	if err2 != nil {
		log.Fatalf("could not get answer: %v", err2)
	}
	log.Println("Token and expires_at are:", resp.Token, resp.ExpiresAt)
}

func main() {
	fmt.Println("Старт gRPC клиента")

	cli = gr.Grpc_connect()
	fmt.Println("Старт клиента")

	r := mux.NewRouter()
	r.HandleFunc("/login", login)
	//r.HandleFunc("/", hello)
	log.Fatal(http.ListenAndServe(":3000", r))

}
