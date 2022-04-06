package models

import (
	"context"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v4"
	"gopkg.in/h2non/gentleman.v2/plugins/query"
	"strings"
	"time"
)

type User struct {
	ID 				uuid.UUID	`json:"id"`
	CreateAt 		time.Time	`json:"_"`
	UpdateAt		time.Time	`json:"_"`
	Email 			string 		`json:"email"`
	PasswordHash 	string	 	`json:"_"`
	Password 		string		`json:"password"`
	PasswordConfirm string 		`json:"password_confirm"`
}

func (u *User) Register(conn *pgx.Conn) error {
	// validate password
	if len(u.Password)<6 {
		return fmt.Errorf("Password must be at least 6 characters long.")
	}
	if u.Password != u.PasswordConfirm {
		return fmt.Errorf("Password do not match.")
	}
	if len(u.Email) < 4 {
		return fmt.Errorf("Incorrect email")
	}

	// validate email
	u.Email = strings.ToLower(u.Email)
	row := conn.QueryRow(context.Background(), "SELECT id from user_account WHERE email=$1", u.Email)
	userLookup := User{}
	err := row.Scan(&userLookup)
	if err != pgx.ErrNoRows {
		fmt.Println("Found user " + userLookup.Email)
		return fmt.Errorf("User with this email alredy exists")
	}

	// create Hash password
	pwdHash, err := bcrypt
}