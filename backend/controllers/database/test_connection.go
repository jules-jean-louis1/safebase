package controllers

import (
	utils "backend/controllers/utils"

	"github.com/gin-gonic/gin"
)

func TestConnection(c *gin.Context) {
	host := c.Query("host")
	port := c.Query("port")
	username := c.Query("username")
	password := c.Query("password")
	dbName := c.Query("dbName")
	dbType := c.Query("dbType")

	// TODO: If intel for the database is the same a the one in the .env file, use the .env file, because we are trying to connect to the database in the container
	if host == "localhost" {
		host = "safebase_db"
	}

	params := &utils.DBParams{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
		DBName:   dbName,
		SSLMode:  "disable",
		DBType:   dbType,
	}

	co, err := utils.ConnectionTester(params)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(200, co)
	}
}
