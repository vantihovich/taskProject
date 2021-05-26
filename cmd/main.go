package main

import (
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	gr "github.com/vantihovich/taskProject/api"
	cnfg "github.com/vantihovich/taskProject/configuration"
	hndl "github.com/vantihovich/taskProject/handlers"
	postgr "github.com/vantihovich/taskProject/postgres"
)

func main() {
	log.WithFields(log.Fields{}).Info("Configs loading")

	cfg, err := cnfg.Load()
	if err != nil {
		log.WithFields(log.Fields{}).Panic("Failed to load app config")
	}

	log.WithFields(log.Fields{}).Info("Connecting to DB")

	db := postgr.New(cfg)

	if err := db.Open(); err != nil {
		log.WithFields(log.Fields{}).Panic("Failed to establish DB connection")
	}

	log.WithFields(log.Fields{}).Info("Starting gRPC client")
	gr.GrpcCliConn()

	log.WithFields(log.Fields{}).Info("Strating client")

	r := mux.NewRouter()
	r.HandleFunc("/login", hndl.Login)
	r.HandleFunc("/", hndl.Hello)

	log.Fatal(http.ListenAndServe(":3000", r))

}
