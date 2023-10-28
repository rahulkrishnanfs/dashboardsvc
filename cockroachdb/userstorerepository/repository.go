package userrepository

import (
	"dashboardsvc/model"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Repository struct {
	db *sql.DB
}

func NewRepository() *Repository {
	dbStore, err := sql.Open("postgres", "postgresql://max:roach@localhost:26257/bank?sslmode=verify-full&sslrootcert=/Users/rahulra/Documents/go/Preparationforsucess/cockroachlabs/certs/ca.crt")
	if err != nil {
		return nil
	}
	return &Repository{
		db: dbStore,
	}
}

func (repo *Repository) Create(user model.User) {
	fmt.Println("esecuting query")
	fmt.Println(user)
	sql := `INSERT INTO users (username, password, firstname, lastname, address, email) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := repo.db.Exec(sql, user.Username, user.Password, user.Firstname, user.Lastname, user.Address, user.Email)
	if err != nil {
		fmt.Println(err)
	}

}
func (repo *Repository) SignIn(username, password string) bool {
	var dbpassword string
	sql := ``
	fmt.Println(sql)
	err := repo.db.QueryRow("select password from users where username='rahul2'").Scan(&dbpassword)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Paswrd validated", dbpassword)
	return password == dbpassword
}
