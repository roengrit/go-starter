package main

import (
	"gm-startd/app"
	userDb "gm-startd/database/users"
	_ "gm-startd/docs"
	swaggerRouter "gm-startd/routes/docs"
	healthRouter "gm-startd/routes/health"
	userRouter "gm-startd/routes/users"
	userUseCase "gm-startd/usecases/user"

	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// @title GM Startd API
// @version 1.0
// @description GM Startd API documentation
// @host localhost:8080
// @BasePath /
func main() {

	// Initialize repository
	rx, ex := app.Init()
	if ex != nil {
		log.Fatalf(`main.Init: %s`, ex.Error())
	}
	time.Local = rx.Location

	userDB, ex := userDb.Init()
	if ex != nil {
		log.Fatalf(`main.userDB: %s`, ex.Error())
	}

	// Initialize use case
	userUC := userUseCase.Init(userDB)

	// Set up router
	router := gin.Default()
	swaggerRouter.Router(router)
	userRouter.Router(userUC, router)
	healthRouter.Router(router)

	// Run the server
	router.Run(":8080")
}
