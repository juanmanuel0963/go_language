package DependencyInjection

// repo.go
type DBInterface interface {
	Get(id string) (string, error)
}

//----------Database----------
type Database struct {
	Host string
	Port int
}

func (r *Database) Get(id string) (string, error) {
	// Implements DB interface
	return id, nil
}

//----------Repository----------
type Repository struct {
	db DBInterface
}

func (r *Repository) Get(id string) (string, error) {
	return r.db.Get(id)
}

func NewRepository(db DBInterface) *Repository {
	return &Repository{db: db}
}
