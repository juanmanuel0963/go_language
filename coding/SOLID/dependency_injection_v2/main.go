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

// ----------DBConn----------
type DBConn interface {
	Query() interface{}
}

// ----------MySQL----------
type MySQL struct {
}

func (db MySQL) Query() interface{} {
	return []string{"alex", "john", "mike"}
}

// ----------PostgreSQL----------
type PostgreSQL struct {
}

func (db PostgreSQL) Query() interface{} {
	return map[string]string{
		"a3f69c2b-d153-48fd-b10c-5b641657477b": "alex",
		"a4f69c2b-d153-48fd-b10c-5b641657477a": "john",
		"a5f69c2b-d153-48fd-b10c-5b641657477c": "mike",
	}
}

// ----------UsersRepository----------
type UsersRepository struct {
	db DBConn
}

func (r UsersRepository) GetUsers() []string {
	var users []string
	res := r.db.Query()

	switch res := res.(type) {
	case map[string]string:
		for _, u := range res {
			users = append(users, u)
		}
		return users
	case []string:
		return res
	}
	return []string{}
}

// ----------main----------
func main() {
	mysqlDB := MySQL{}
	postgreSQLDB := PostgreSQL{}
	repo1 := UsersRepository{db: mysqlDB}
	repo2 := UsersRepository{db: postgreSQLDB}
	fmt.Println(repo1.GetUsers())
	fmt.Println(repo2.GetUsers())
}
