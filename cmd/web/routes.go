package main

import (
	"github.com/gin-gonic/gin"
)

// Routes for user object
func (app *App) userRoutes(router *gin.Engine) {
	router.GET("/users", app.getUsers())
	router.GET("/users/:user_id", app.getUser())
	router.POST("/user/signup", app.signup())
	router.POST("/user/login", app.login())

}

// Routes for food item
func (app *App) foodRoutes(router *gin.Engine) {
	router.GET("/foods", app.getFoods())
	router.GET("/foods/:food_id", app.getFood())
	router.POST("/foods", app.createFood())
	router.PATCH("/foods/:food_id", app.updateFood())
}

// routes for invoice
func (app *App) invoiceRoutes(router *gin.Engine) {
	router.GET("/invoices", app.getInvoices())
	router.GET("/invoices/invoice_id", app.getInvoice())
	router.POST("/invoices", app.createInvoice())
	router.PATCH("/invoices/invoice_id", app.updateInvoice())
}

// routes for menu item
func (app *App) menuRoutes(router *gin.Engine) {
	router.GET("/menus", app.getMenus())
	router.GET("/menus/menu_id", app.getMenu())
	router.POST("/menus", app.createMenu())
	router.PATCH("/menus/menu_id", app.updateMenu())
}

// routes for order placed by user
func (app *App) orderRoutes(router *gin.Engine) {
	router.GET("/orders", app.getOrders())
	router.GET("/orders/order_id", app.getOrder())
	router.POST("/order", app.createOrder())
	router.PATCH("/order/order_id", app.updateOrder())
}

// routes for a particular item ordered
func (app *App) orderItemRoutes(router *gin.Engine) {
	router.GET("/items", app.getItems())
	router.GET("/items/item_id", app.getItem())
	router.GET("/items-order/:order_id", app.getItemsByOrder())
	router.POST("/items", app.createItem())
	router.PATCH("/items/item_id", app.updateItem())
}

// routes for table object
func (app *App) tableRoutes(router *gin.Engine) {
	router.GET("/tables", app.getTables())
	router.GET("/tables/table_id", app.getTable())
	router.POST("/tables", app.createTable())
	router.PATCH("/tables/table_id", app.updeateTable())
}
