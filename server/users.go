package main

import (
	"fmt"

	gr "github.com/vantihovich/taskProject/api"
)

func main() {
	fmt.Println("Старт сервера")

	gr.GrpcServConn()

}
