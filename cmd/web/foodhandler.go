package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Zekeriyyah/restaurant-api/internals/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (app *App) getFoods() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func (app *App) getFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		food := &models.Food{}

		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		foodId := c.Param("food_id")
		err := foodCollection.FindOne(ctx, bson.M{"food_id": foodId}).Decode(food)
		if err != nil {
			app.errorLog.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while fetching the food items"})
		}
		c.JSON(http.StatusOK, food)
	}
}

func (app *App) createFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var menu models.Menu
		var food models.Food

		if err := c.BindJSON(&food); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		ERR_VALIDATION := validate.Struct(food)
		if ERR_VALIDATION != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": ERR_VALIDATION.error()})
		}
		err := menuCollection.FindOne(ctx, bson.M{"menu_id": food.Menu_id}).Decode(&menu)
		if err != nil {
			msg := fmt.Sprintf("menu not found")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		food.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		food.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		food.ID = primitive.NewObjectID()
		food.Food_id = food.ID.Hex()
		
		val := app.toFixed(*food.Price, 2)
		food.Price = &val

		result, ERR_INSERT := foodCollection.InsertOne(ctx,food)		
		if ERR_INSERT != nil {
			msg := fmt.Sprintf("food item not created...")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		}

		c.JSON(http.StatusOK, result)
}

func (app *App) updateFood() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func (app *App) toFixed(num float64, precision string) float64 {

}
