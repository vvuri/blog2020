package expgin

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/h2non/baloo.v3"
	"testing"
)

// https://github.com/gin-gonic/gin
// https://github.com/h2non/baloo

// go get -u github.com/gin-gonic/gin
// go get -u gopkg.in/h2non/baloo.v3
// go mod vendor

func runserver() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}

var tests = baloo.New("http://127.0.0.1:8080")

// http://127.0.0.1:8080/ping -> {"message":"pong"}
func TestSimpleGin(t *testing.T) {
	go runserver()

	tests.Get("/ping").
		Expect(t).
		Status(200).
		Header("Content-Type", "application/json; charset=utf-8").
		Type("json").
		JSON(map[string]string{"message": "pong"}).
		Done()
}

