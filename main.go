package main

import (
	"github.com/IvanoffDan/event-sourcing/routers"
)

func main() {
	r := routers.InitRoutes() // Initialize gin router
	r.Run()                   // listen and serve on 0.0.0.0:8080
}
