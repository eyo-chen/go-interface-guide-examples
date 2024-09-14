package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/eyo-chen/a-straignforward-guid-for-go-interface/example04-decoupling/with-interface/handler"
	"github.com/eyo-chen/a-straignforward-guid-for-go-interface/example04-decoupling/without-interface/service"
)

func main() {
	conn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/test")
	if err != nil {
		log.Fatal(err)
	}

	service := service.NewMySQLService(conn)
	handler := handler.NewHandler(service)

	http.HandleFunc("/user", handler.GetUser)
}
