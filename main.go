package main

import (
	"IM_project/router" // router "ginchat/router"
)

func main() {
	r := router.Router() // router.Router()
	r.Run(":8081")       // listen and serve on 0.0.0.0:8080 (forwindows "localhost:8080")
}
