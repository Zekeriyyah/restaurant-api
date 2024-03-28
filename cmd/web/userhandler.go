package main

import "github.com/gin-gonic/gin"

func (app *App) getUsers() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func (app *App) getUser() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
func (app *App) signup() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
func (app *App) login() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func (app *App) hashPasswd(passwd string) string {
	return ""
}

func (app *App) verifyPasswd(validPasswd, passwd string) (bool, string) {
	return true, ""
}
