package routes

import (
	"github.com/W5DEV/go-recipe-api/controllers"
	"github.com/W5DEV/go-recipe-api/middleware"
	"github.com/gin-gonic/gin"
)

type RecipeRouteController struct {
	recipeController controllers.RecipeController
}

func NewRouteRecipeController(recipeController controllers.RecipeController) RecipeRouteController {
	return RecipeRouteController{recipeController}
}

func (pc *RecipeRouteController) RecipeRoute(rg *gin.RouterGroup) {

	router := rg.Group("recipes")
	router.Use(middleware.DeserializeUser())
	router.POST("/", pc.recipeController.CreateRecipe)
	router.GET("/", pc.recipeController.FindRecipes)
	router.PUT("/:recipeId", pc.recipeController.UpdateRecipe)
	router.GET("/:recipeId", pc.recipeController.FindRecipeById)
	router.DELETE("/:recipeId", pc.recipeController.DeleteRecipe)
}