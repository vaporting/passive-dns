package main

import (
	"passive-dns/routes"

	"fmt"

	"passive-dns/db"
)

func main() {
	db.InitDB()
	_, err := db.GetDB()
	fmt.Println(err)
	r := routes.InitRoutes()
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
