package main

import (
	log "github.com/sirupsen/logrus"
	gr "github.com/vantihovich/taskProject/api"
)

func main() {

	log.WithFields(log.Fields{}).Info("Server starting")

	gr.GrpcServConn()

}
