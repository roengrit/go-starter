package entities

// data definitions
type User struct {
	ID   int    `json:"id"`   // UserId
	Name string `json:"name"` // UserName
}

type Users struct {
	Data *[]User `json:"data"` // Users
}
