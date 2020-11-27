package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/vvuri/blog2020/cmd/goblog/docs"
	//"./docs"
)

var db = make(map[string]string)

// @title Goblog Swagger API
// @version 1.0
// @description Swagger API for Golang Project Goblog.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email vvuri1978@gmail.com
// @license.name MIT
// @BasePath /api/v1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func creteSwaggerDoc() {
	docs.SwaggerInfo.Title = "Swagger API"
	docs.SwaggerInfo.Description = "Goblog RRST API server."
	docs.SwaggerInfo.Host = "localhost:5000"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "api/v1"
	docs.SwaggerInfo.Schemes = []string{"http"}
}

func setupRouter() *gin.Engine {
	// Disable Console Color - gin.DisableConsoleColor()
	// Создать Router
	r := gin.Default()

	// Add for Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Добавить обработчик - Ping test
	r.GET("/ping", GetPing)

	// Get user value
	r.GET("/user/:name", GetName)

	// Authorized group (uses gin.BasicAuth() middleware)
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	authorized.POST("admin", PostAdmin)

	return r
}

// GetName godoc
// @Summary Retrieves user based on given ID
// @Produce json
// @Param id path integer true "User ID"
// @Success 200 {object} models.User
// @Router /user/:name [get]
// @Security ApiKeyAuth
func GetName(c *gin.Context) {
	user := c.Params.ByName("name")
	value, ok := db[user]
	if ok {
		c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
	} else {
		c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
	}
}

func GetPing(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func PostAdmin(c *gin.Context) {
	user := c.MustGet(gin.AuthUserKey).(string)

	// Parse JSON
	var json struct {
		Value string `json:"value" binding:"required"`
	}

	if c.Bind(&json) == nil {
		db[user] = json.Value
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	}
}

func main() {
	creteSwaggerDoc()

	r := setupRouter()
	err := r.Run("127.0.0.1:5000")

	if err != nil {
		panic(err)
	}
}
