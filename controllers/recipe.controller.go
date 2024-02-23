package controllers

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/W5DEV/go-recipe-api/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RecipeController struct {
	DB *gorm.DB
}

func NewRecipeController(DB *gorm.DB) RecipeController {
	return RecipeController{DB}
}

// Create Recipe Handler
func (pc *RecipeController) CreateRecipe(ctx *gin.Context) {
	var payload *models.CreateRecipeRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	now := time.Now()
	newRecipe := models.Recipe{
		Title:         payload.Title,
		Description:   payload.Description,
		Ingredients:   payload.Ingredients,
		Instructions:  payload.Instructions,
		Image:         payload.Image,
		Chef:          payload.Chef,
		Inactive:	   payload.Inactive,
		CreatedAt:     now,
		UpdatedAt:     now,
	}

	result := pc.DB.Create(&newRecipe)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "Recipe with that title already exists"})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newRecipe})
}

// Update Recipe Handler
func (pc *RecipeController) UpdateRecipe(ctx *gin.Context) {
	recipeId := ctx.Param("recipeId")

	var payload *models.UpdateRecipe
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	var updatedRecipe models.Recipe
	result := pc.DB.First(&updatedRecipe, "id = ?", recipeId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No recipe with that title exists"})
		return
	}
	now := time.Now()
	recipeToUpdate := models.Recipe{
		Title:         payload.Title,
		Description:   payload.Description,
		Ingredients:   payload.Ingredients,
		Instructions:  payload.Instructions,
		Image:         payload.Image,
		Chef:          payload.Chef,
		Inactive:	   payload.Inactive,
		CreatedAt:     now,
		UpdatedAt:     now,
	}

	pc.DB.Model(&updatedRecipe).Updates(recipeToUpdate)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": updatedRecipe})
}

// Get Single Recipe Handler
func (pc *RecipeController) FindRecipeById(ctx *gin.Context) {
	recipeId := ctx.Param("recipeId")

	var recipe models.Recipe
	result := pc.DB.First(&recipe, "id = ?", recipeId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No recipe with that title exists"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": recipe})
}

// Get All Recipe Handler
func (pc *RecipeController) FindRecipes(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var recipes []models.Recipe
	results := pc.DB.Limit(intLimit).Offset(offset).Find(&recipes)
	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(recipes), "data": recipes})
}

// Delete Recipe Handler
func (pc *RecipeController) DeleteRecipe(ctx *gin.Context) {
	recipeId := ctx.Param("recipeId")

	result := pc.DB.Delete(&models.Recipe{}, "id = ?", recipeId)

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No recipe with that title exists"})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

