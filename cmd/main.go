package main

import (
	"fmt"

	"log"
	"net/http"

	"github.com/gorilla/mux"
	gr "github.com/vantihovich/taskProject/api"
	cnfg "github.com/vantihovich/taskProject/configuration"
	hndl "github.com/vantihovich/taskProject/handlers"
	postgr "github.com/vantihovich/taskProject/postgres"
)

func main() {
	fmt.Println("Configs loading")
	cfg, err := cnfg.Load()
	if err != nil {
		panic("Failed to load app config")
	}
	fmt.Println("Connecting to DB")

	db := postgr.New(cfg)

	if err := db.Open(); err != nil {
		panic("Failed to establish DB connection")
	}

	fmt.Println("Starting gRPC client")
	gr.GrpcCliConn()

	fmt.Println("Strating client")

	r := mux.NewRouter()
	r.HandleFunc("/login", hndl.Login)
	r.HandleFunc("/", hndl.Hello)

	log.Fatal(http.ListenAndServe(":3000", r))

}
