package queries

import (
	"github.com/create-go-app/fiber-go-template/app/models"
	"github.com/jmoiron/sqlx"
)

type InstanceQueries struct {
	*sqlx.DB
}

// CreateUser query for creating a new instance
func (q *InstanceQueries) CreateInstance(u *models.GCEInstance) error {
	// Define query string.
	query := `INSERT INTO instances VALUES ($1, $2, $3, $4, $5, $6, $7)`

	// Send query to database.
	_, err := q.Exec(
		query,
		u.ID, u.CreatedAt, u.UpdatedAt, u.Name, u.Zone, u.Status, u.Compute.ID,
	)
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}

// GetInstances method for getting all instances.
func (q *InstanceQueries) GetInstances() ([]models.GCEInstance, error) {
	// Define instance collection.
	instances := []models.GCEInstance{}

	// Define query string.
	query := `SELECT * FROM instances`

	// Send query to database.
	err := q.Select(&instances, query)
	if err != nil {
		// Return empty object and error.
		return instances, err
	}

	// Return query result.
	return instances, nil
}