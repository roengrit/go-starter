package entities

import "time"

// data definitions
type User struct {
	ID           int        `json:"id"`                      // UserId
	Username     string     `json:"username"`                // UserName
	Name         string     `json:"name"`                    // Name
	NickName     string     `json:"nick_name"`               // NickName
	Email        string     `json:"email"`                   // Email
	PasswordHash string     `json:"password_hash,omitempty"` //PasswordHash
	CreatedAt    time.Time  `json:"created_at"`              //CreatedAt
	UpdatedAt    time.Time  `json:"updated_at"`              //UpdatedAt
	LastLogin    *time.Time `json:"last_login"`              // LastLogin

}

type Users struct {
	Data *[]User `json:"data"` // Users
}
