package main

import (
	"log"
	"os"

	"github.com/Zekeriyyah/restaurant-api/internals"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	infoLog  *log.Logger
	errorLog *log.Logger
	version  string
}

var ( 
	dbClient, ERR_CLIENT = internals.DBInstance()
	foodCollection *mongo.Collection = internals.OpenCollection(dbClient, "food")
	menuCollection *mongo.Collection = internals.OpenCollection(dbClient, "menu")
	validate = validator.New()
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &App{
		infoLog:  infoLog,
		errorLog: errorLog,
		version:  "1.0",
	}

	//Setting up database connection
	if ERR_CLIENT == nil {
		app.infoLog.Println("database connection successful....")
	}

	//set-up the router
	mux := gin.New()
	mux.Use(gin.Logger())
	//app.userRoutes(mux)
	//mux.Use(app.userAuth())

	//passing router to each routes
	app.foodRoutes(mux)
	//app.invoiceRoutes(mux)
	//app.menuRoutes(mux)
	//app.orderRoutes(mux)
	//	app.orderItemRoutes(mux)
	//app.tableRoutes(mux)

	mux.Run(":" + port)
}
