package repository

import (
	"database/sql"

	"github.com/akanshgupta98/BlogProject/user-service/internal/database"
)

type Repo struct {
	db *sql.DB
}

type UserTable struct {
	ID       int
	Username string
	Email    string
	Name     string
	Phone    string
}

func RepoInit() (*Repo, error) {
	dbHdlr, err := database.NewDB()
	if err != nil {
		return nil, err
	}
	return &Repo{
		db: dbHdlr,
	}, nil
}

func (r *Repo) CreateUser(data UserTable) (int, error) {
	rows, err := r.db.Query("INSERT INTO users (username,name,email,phone) VALUES ($1,$2,$3,$4) RETURNING id", data.Username, data.Name, data.Email, data.Phone)
	if err != nil {
		return -1, err
	}
	var result int

	defer rows.Close()
	for rows.Next() {
		var id int

		err := rows.Scan(&id)
		if err != nil {
			return -1, err
		}
		result = id
	}

	return result, nil

}

func (r *Repo) FetchUserByID(data UserTable) (UserTable, error) {
	tmp := UserTable{}
	rows, err := r.db.Query("SELECT id,username,name,email,phone FROM users WHERE id= $1 ", data.ID)
	if err != nil {
		return tmp, err
	}

	defer rows.Close()
	for rows.Next() {

		err := rows.Scan(&tmp.ID, &tmp.Username, &tmp.Name, &tmp.Email, &tmp.Phone)
		if err != nil {
			return tmp, err
		}

	}

	return tmp, nil

}
