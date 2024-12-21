package main

import (
	"awesomeProject12/Controller"
	"awesomeProject12/initializers"
	"github.com/gin-gonic/gin"
)

// this func is before an server run so like
// it can be .env connect to database ect
func init() {
	//loading env file
	initializers.LoadEnv()
	//connect to db
	initializers.ConnectDB()
	//create table
	initializers.CreateDb()

}

// main func where we have all methods and we run server
func main() {
	//creating server default
	r := gin.Default()
	//all the methods
	//GET METHOD
	//GET ALL USERS
	r.GET("/", Controller.GetAllUsers)
	//GET USER BY ID
	r.GET("/:id", Controller.GetUserById)
	//GET USERS BY SEARCH
	r.GET("/search", Controller.GetUserBySearch)
	//POST METHOD
	//ADDING USER
	r.POST("/", Controller.AddUser)
	//PUT METHOD
	//update user by id
	r.PUT("/:id", Controller.UpdateUserById)
	//DELETE METHOD
	//delete user by id
	r.DELETE("/:id", Controller.DeleteUserById)
	//running the server
	r.Run(":8080")
}
