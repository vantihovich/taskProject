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
	cnfg "github.com/vantihovich/taskProject/configuration"
	postgr "github.com/vantihovich/taskProject/postgres"
)

func main() {
	fmt.Println("Загрузга конфигов")
	cfg, err := cnfg.Load()
	if err != nil {
		panic("Failed to load app config")
	}
	fmt.Println("Установка связи с БД")

	db := postgr.New(cfg)

	if err := db.Open(); err != nil {
		panic("Failed to establish DB connection")
	}

	fmt.Println("Старт gRPC клиента")
	gr.GrpcCliConn()

	fmt.Println("Старт клиента")

	r := mux.NewRouter()
	r.HandleFunc("/login", hndl.Login)
	r.HandleFunc("/", hndl.Hello)

	log.Fatal(http.ListenAndServe(":3000", r))

}
