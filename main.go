package main

import "stori/routes"

// @title API stori
// @version 0.0.1
// @description System that processes transactions

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 52.202.149.44
// @BasePath /api/stori
func main() {
	router := routes.NewRouter()

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
