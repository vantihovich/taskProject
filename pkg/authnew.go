package main

import (
	"fmt"

	"log"
	"net/http"

	"github.com/gorilla/mux"
	gr "github.com/vantihovich/taskProject/api"

	//dbconn "github.com/vantihovich/taskProject/dbConn"
	hndl "github.com/vantihovich/taskProject/handlers"
	//internal "github.com/vantihovich/taskProject/internal"
	postgr "github.com/vantihovich/taskProject/postgres"
)

func main() {
	fmt.Println("Старт gRPC клиента")

	gr.GrpcCliConn()

	fmt.Println("Установка связи с БД")

	postgr.DB.Open()

	//dbconn.DbConn()

	fmt.Println("Старт клиента")

	r := mux.NewRouter()
	r.HandleFunc("/login", hndl.Login)
	r.HandleFunc("/", hndl.Hello)

	log.Fatal(http.ListenAndServe(":3000", r))

}
