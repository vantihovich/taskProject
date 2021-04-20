package main

import (
	"encoding/json"
	"fmt"

	"io"

	//"strings"

	"log"
	"net/http"

	"github.com/gorilla/mux"
	ps "github.com/vantihovich/taskProject/proto"
	"google.golang.org/grpc"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Эта строка должна быть видна в браузере")
}

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

}

func main() {
	fmt.Println("Старт сервиса")

	r := mux.NewRouter()

	r.HandleFunc("/login", login)
	r.HandleFunc("/", hello)

	log.Fatal(http.ListenAndServe(":3000", r))

	fmt.Println("Старт gRPC")

	conn, err := grpc.Dial("127.0.0.1:3500", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	client := ps.NewTokenGeneratorServiceClient(conn)

	resp, err := client.GenerateToken(context.Background(),
		&ps.Request{params})

	if err != nil {
		log.Fatalf("could not get answer: %v", err)
	}
	log.Println("Token and expires_at are:", resp.token, resp.expires_at)
}
