package expgin

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"github.com/vvuri/blog2020/cmd/expgin/routes"
	"log"
	"testing"
)

// docker pull postgres
// docker run --name gopostgres -e POSTGRES_PASSWORD=rowssap -p 5432:5432 -d postgres
// docker exec -it gopostgres psql -U postgres

// psql -> pgAdmin -> create DB "offersapp"
// $ psql offersapp < database/schema.sql

// Post to localhost:3000/users/register  -> { "user_id": "someid" }


func runserver() {
	conn, err := connectDB()
	if err!=nil {
		return
	}

	log.Println(conn)

	router := gin.Default()

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

func TestSecondRun(t *testing.T) {
	//go
	runserver()
}
