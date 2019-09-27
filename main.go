package main

import "passive-dns/routes"

func main() {
	r := routes.InitRoutes()
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
