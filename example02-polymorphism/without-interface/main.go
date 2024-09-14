package main

import (
	"database/sql"
	"fmt"

	"github.com/eyo-chen/a-straignforward-guid-for-go-interface/example02-polymorphism/without-interface/myorm"
)

func main() {
	conn, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		// panic(err)
		fmt.Println("err")
	}

	user := map[string]interface{}{
		"id":   1,
		"name": "John Doe",
	}
	orm := myorm.New(myorm.MySQL{Conn: conn})
	orm.Insert("users", user)
}
