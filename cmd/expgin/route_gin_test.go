package expgin

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/h2non/baloo.v3"
	"testing"
)

const PORT = ":10001"
var test = baloo.New("http://127.0.0.1"+PORT)

func addRoute() {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	router.GET("/someGet", getting)
	router.POST("/somePost", posting)
	router.PUT("/somePut", putting)
	//router.DELETE("/someDelete", deleting)
	//router.PATCH("/somePatch", patching)
	//router.HEAD("/someHead", head)
	router.OPTIONS("/someOptions", options)

	router.Run(PORT)
}

func getting(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "someGet",
	})
}

func posting(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "somePost",
	})
}

func putting(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "somePut",
	})
}

func options(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "someOptions",
	})
}

func TestRouteGet(t *testing.T) {
	go addRoute()

	test.Get("/someGet").
		Expect(t).
		Status(200).
		Type("json").
		JSON(map[string]string{"message": "someGet"}).
		Done()
}

func TestRoutePost(t *testing.T) {
	go addRoute()

	test.Post("/somePost").
		Expect(t).
		Status(200).
		Type("json").
		JSON(map[string]string{"message": "somePost"}).
		Done()
}

func TestRoutePut(t *testing.T) {
	go addRoute()

	test.Put("/somePut").
		Expect(t).
		Status(200).
		Type("json").
		JSON(map[string]string{"message": "somePut"}).
		Done()
}

func TestRouteOptions(t *testing.T) {
	go addRoute()

	test.Method("/someOptions").
		SetHeader("Send", "Options")
}