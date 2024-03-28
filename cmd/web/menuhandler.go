package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Zekeriyyah/restaurant-api/internals/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

func (app *App) getMenus() gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		result, err := menuCollection.Find(context.TODO(), bson.M{})

		if err != nil {
			msg := fmt.Sprintf("error occured while listing the menu items")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		}
		var allMenus []bson.M
		if err = result.All(ctx, &allMenus); err != nil {
			app.errorLog.Fatal(err)
		}
		c.JSON(http.StatusOK, allMenus)

	}
}

func (app *App) getMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var menu models.Menu
		menuId := c.Param("menu_id")

		err := menuCollection.FindOne(ctx, bson.M{"menu_id": menuId}).Decode(&menu)
		if err != nil {
			msg := fmt.Sprintf("error occured while fetching the menu")
			c.JSON(http.StatusBadRequest, gin.H{"error": msg})
		}
		c.JSON(http.StatusOK, menu)
	}
}

func (app *App) createMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		var menu models.Menu

		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		err := c.BindJSON(&menu)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ERR_VALIDATION := validate.Struct(&menu)
		if ERR_VALIDATION != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": ERR_VALIDATION.Error()})
			return
		}

		menu.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		menu.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		menu.ID = primitive.NewObjectID()
		menu.Menu_id = menu.ID.Hex()

		result, ERR_INSERT := menuCollection.InsertOne(ctx, menu)
		if ERR_INSERT != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": ERR_INSERT.Error()})
			return
		}

		c.JSON(http.StatusOK, result)
	}
}

func (app *App) inTimeSpan(start, end, t time.Time) bool {
	return start.After(t) && end.After(start)
}

func (app *App) updateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		var menu models.Menu
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		err := c.BindJSON(menu)
		if err != nil {
			app.errorLog.Println("error occured while binding JSON request body to menu object!")
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		menuId := c.Param("menu_id")
		filter := bson.M{"menu_id": menuId}

		var updateObj primitive.D

		if menu.Start_Date != nil && menu.End_Date != nil {
			if !app.inTimeSpan(*menu.Start_Date, *menu.End_Date, time.Now()) {
				msg := "please retype the time appropriately"
				c.JSON(http.StatusBadRequest, gin.H{"error": msg})
				return
			}

			updateObj = append(updateObj, bson.E{"start_date", menu.Start_Date})
			updateObj = append(updateObj, bson.E{"end_date", menu.End_Date})

			if menu.Name != "" {
				updateObj = append(updateObj, bson.E{"name", menu.Name})
			}

			if menu.Category != "" {
				updateObj = append(updateObj, bson.E{"category", menu.Category})
			}

			menu.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
			updateObj = append(updateObj, bson.E{"updated_at", menu.Updated_at})

			upsert := true
			opt := options.UpdateOptions{
				Upsert: &upsert,
			}

			result, err := menuCollection.UpdateOne(
				ctx, 
				filter, 
				bson.D{
					{"$set", updateObj}
				}, 
				&opt,
			)

			if err != nil {
				msg := "menu update failed!"
				c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			}

			c.JSON(http.StatusOK, result) 
		}

	}
}
