package controllers

import (
	utils "backend/controllers/utils"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TestConnection(c *gin.Context) {
	host := c.Query("host")
	port := c.Query("port")
	username := c.Query("username")
	password := c.Query("password")
	dbName := c.Query("dbName")
	dbType := c.Query("dbType")

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

type DBParams struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

func TestF(c *gin.Context) {
	params := &DBParams{
		Host:     "host.docker.internal",
		Port:     "3306",
		Username: "jj",
		Password: "password",
		DBName:   "silver_micro",
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		params.Username, params.Password, params.Host, params.Port, params.DBName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Connection successful"})
}
