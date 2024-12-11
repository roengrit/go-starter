package users

import (
	"database/sql"
	"gm-startd/entities"
	"os"

	_ "github.com/lib/pq"
)

type UserDB struct {
	db *sql.DB
}

func Init() (*UserDB, error) {
	connectionString := os.Getenv("DATABASE_USER_DNS")
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	return &UserDB{db: db}, nil
}

func (repo *UserDB) GetAll() (*entities.Users, error) {
	rows, err := repo.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []entities.User
	for rows.Next() {
		var user entities.User
		err := rows.Scan(&user.ID, &user.Name)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	data := &entities.Users{Data: &users}
	return data, nil
}

func (repo *UserDB) GetByID(id int) (*entities.User, error) {
	var user entities.User
	err := repo.db.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *UserDB) CreateNew(user *entities.User) error {
	_, err := repo.db.Exec("INSERT INTO users (id, name) VALUES ($1, $2)", user.ID, user.Name)
	return err
}

func (repo *UserDB) DeleteByID(id int) error {
	_, err := repo.db.Exec("DELETE FROM users WHERE id = $1", id)
	return err
}
