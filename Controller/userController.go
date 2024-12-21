package Controller

import (
	"awesomeProject12/initializers"
	"awesomeProject12/models"
	"crypto/sha256"
	"fmt"
	"github.com/gin-gonic/gin"
)

// GET
// all users
func GetAllUsers(c *gin.Context) {
	var users []models.User
	initializers.DB.Find(&users)
	c.JSON(200, gin.H{"users": users})
}

// by id
func GetUserById(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	// Use First to find the user by ID
	if err := initializers.DB.First(&user, id).Error; err != nil {
		// Handle case where user is not found
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	// Return the found user
	c.JSON(200, gin.H{"user": user})
}

// by search
func GetUserBySearch(c *gin.Context) {
	var users []models.User
	value := c.Query("search")

	initializers.DB.Where("name LIKE ? OR last_name LIKE ? OR email LIKE ? OR telephone LIKE ?",
		"%"+value+"%", "%"+value+"%", "%"+value+"%", "%"+value+"%").Find(&users)
	c.JSON(200, gin.H{"users": users})
}

// POST
// adding user we need like pass need to be 1-190 char  email ect
func AddUser(c *gin.Context) {
	// Create a variable to hold user data from the request
	var user models.User

	// Bind JSON data from the request body to the user struct
	if err := c.ShouldBindJSON(&user); err != nil {
		// Return error if JSON binding fails
		c.JSON(400, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	// Hash the password using SHA256
	hashedPassword := sha256.Sum256([]byte(user.Password))
	user.Password = fmt.Sprintf("%x", hashedPassword)

	// Save the user to the database
	if err := initializers.DB.Create(&user).Error; err != nil {
		// Handle any error that might occur when saving the user in the database
		c.JSON(500, gin.H{"error": "Failed to create user: " + err.Error()})
		return
	}

	// Return success status and the created user
	c.JSON(201, gin.H{"message": "User created successfully", "user": user})
}

// PUT
// we need to search users
// we need by an id
func UpdateUserById(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	// Use First to find the user by ID
	if err := initializers.DB.First(&user, id).Error; err != nil {
		// Handle case where user is not found
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	// Bind JSON data from the request body to the user struct
	if err := c.ShouldBindJSON(&user); err != nil {
		// Return error if JSON binding fails
		c.JSON(400, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	// Update the user details in the database
	if err := initializers.DB.Save(&user).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to update user: " + err.Error()})
		return
	}

	// Return the updated user
	c.JSON(200, gin.H{"message": "User updated successfully", "user": user})
}

// DELETE
// by id
func DeleteUserById(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	// Use First to find the user by ID
	if err := initializers.DB.Delete(&user, id).Error; err != nil {
		// Handle case where user is not found
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	// Return the found user
	c.JSON(200, gin.H{"user": "user by id " + id + " is deleted"})
}
