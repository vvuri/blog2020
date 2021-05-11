package expgin

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"github.com/vvuri/blog2020/cmd/expgin/routes"
	"log"
	"testing"
)

func runserver() {
	conn, err := connectDB()
	if err!=nil {
		return
	}

	log.Println(conn)

	router := gin.Default()
	router.Use(dbMiddleware(*conn))

	userGroup := router.Group("users")
	{
		userGroup.POST("register", routes.UserRegister)
	}

	router.Run(":3000")
}

func connectDB() (c *pgx.Conn, err error) {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	conn, err := pgx.Connect(context.Background(), "postgresql://postgres:rowssap@localhost:5432") // /gopostgres")
	if err!=nil {
		log.Println("Error connect to DB")
		log.Println(err.Error())
	}

	_ = conn.Ping(context.Background())
	return conn, err
}

func dbMiddleware(con pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", con)
		c.Next()
	}
}

func TestSecondRun(t *testing.T) {
	//go
	runserver()
}
