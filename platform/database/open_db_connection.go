package database

import "github.com/create-go-app/fiber-go-template/app/queries"

// Queries struct for collect all app queries.
type Queries struct {
	*queries.UserQueries // load queries from User model
	*queries.ComputeQueries // load queries from Compute model
	*queries.InstanceQueries // load queries from Instance model
}

// OpenDBConnection func for opening database connection.
func OpenDBConnection() (*Queries, error) {
	// Define a new PostgreSQL connection.
	db, err := PostgreSQLConnection()
	if err != nil {
		return nil, err
	}

	return &Queries{
		// Set queries from models:
		UserQueries: &queries.UserQueries{DB: db}, // from User model
		ComputeQueries: &queries.ComputeQueries{DB: db},
		InstanceQueries: &queries.InstanceQueries{DB: db},
	}, nil
}
