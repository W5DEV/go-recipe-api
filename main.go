package main

import (
	"log"
	"net/http"

	"github.com/W5DEV/go-recipe-api/controllers"
	"github.com/W5DEV/go-recipe-api/initializers"
	"github.com/W5DEV/go-recipe-api/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)


var (
	server              *gin.Engine
	AuthController      controllers.AuthController
	AuthRouteController routes.AuthRouteController

	UserController      controllers.UserController
	UserRouteController routes.UserRouteController

	RecipeController      controllers.RecipeController
	RecipeRouteController routes.RecipeRouteController
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)

	AuthController = controllers.NewAuthController(initializers.DB)
	AuthRouteController = routes.NewAuthRouteController(AuthController)

	UserController = controllers.NewUserController(initializers.DB)
	UserRouteController = routes.NewRouteUserController(UserController)

	RecipeController = controllers.NewRecipeController(initializers.DB)
	RecipeRouteController = routes.NewRouteRecipeController(RecipeController)

	server = gin.Default()
}

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowCredentials = true
	corsConfig.AddAllowHeaders("Authorization")


	server.Use(cors.New(corsConfig))

	router := server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		message := "Welcome to the go-recipe-api!"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	AuthRouteController.AuthRoute(router)
	UserRouteController.UserRoute(router)
	RecipeRouteController.RecipeRoute(router)
	log.Fatal(server.Run(":" + config.ServerPort))
}