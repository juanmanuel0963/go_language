package main

import (
	"fmt"
)

/*
•	Coming to the “Dependency Inversion Principle” …
the goal is that there are no concrete class names buried inside the application code.
•	The idea is the code only mentions abstract classes (interfaces)…
•	The concrete implementations are injected or provided by factories.
•	This principle is more focus on coding techniques than high level design
*/

// ----------MySQL----------
type MySQL struct {
}

func (db MySQL) QueryMySQL() []string {
	return []string{"alex", "john", "mike"}
}

// ----------PostgreSQL----------
type PostgreSQL struct {
}

func (db PostgreSQL) QueryPostgreSQL() map[string]string {
	return map[string]string{
		"a3f69c2b-d153-48fd-b10c-5b641657477b": "alex",
		"a4f69c2b-d153-48fd-b10c-5b641657477a": "john",
		"a5f69c2b-d153-48fd-b10c-5b641657477c": "mike",
	}
}

// ----------UsersRepository----------
type UsersRepository struct {
	db MySQL
	//db PostgreSQL
}

func (r UsersRepository) GetUsers() []string {
	res := r.db.QueryMySQL()
	//res := r.db.QueryPostgreSQL()
	//var users []string
	//for _, u := range res {
	//	users = append(users, u)
	//}
	//return users
	return res
}

// ----------main----------
func main() {
	mysqlDB := MySQL{}
	//postgreSQLDB := PostgreSQL{}
	repo := UsersRepository{db: mysqlDB}
	//repo := UsersRepository{db: postgreSQLDB}
	fmt.Println(repo.GetUsers())
}
