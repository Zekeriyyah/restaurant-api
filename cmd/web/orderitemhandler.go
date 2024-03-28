package main

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (app *App) getItems() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func (app *App) getItemsByOrder() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func (app *App) itemsByOrder(id string) ([]primitive.M, error) {
	return nil, nil
}

func (app *App) getItem() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func (app *App) createItem() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func (app *App) updateItem() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
