package models

import (
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v4"
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
	if len(u.Password)<6 {
		return fmt.Errorf("Password must be at least 6 characters long.")
	}

	if u.Password != u.PasswordConfirm {
		return fmt.Errorf("Password do not match.")
	}

	if len(u.Email) < 4 {
		return fmt.Errorf("Incorrect email")
	}

	u.Email = strings.ToLower(u.Email)

}