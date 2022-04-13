package queries

import (
	"github.com/create-go-app/fiber-go-template/app/models"
	"github.com/jmoiron/sqlx"
)

type ComputeQueries struct {
	*sqlx.DB
}

// CreateCompute query for creating a new compute
func (q *ComputeQueries) CreateCompute(c *models.Compute) error {
	// Define query string.
	query := `INSERT INTO computes VALUES ($1, $2, $3, $4, $5, $6, $7)`

	// Send query to database.
	_, err := q.Exec(
		query,
		c.ID, c.CreatedAt, c.UpdatedAt, c.Name, c.JsonPath, c.ProjectID, c.Zone,
	)
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}

// GetComputes method for getting all compute objects.
func (q *ComputeQueries) GetComputes() ([]models.Compute, error) {
	computes := []models.Compute{}

	// Define query string.
	query := `SELECT * FROM computes`

	// Send query to database.
	err := q.Select(&computes, query)
	if err != nil {
		// Return empty object and error.
		return computes, err
	}

	// Return query result.
	return computes, nil
}

//  GetCompute method for getting one compute by given ID.
func (q *ComputeQueries) GetCompute(id int) (models.Compute, error) {
	compute := models.Compute{}

	// Define query string.
	query := `SELECT * FROM computes WHERE id = $1`

	// Send query to database.
	err := q.Get(&compute, query, id)
	if err != nil {
		// Return empty object and error.
		return compute, err
	}

	// Return query result.
	return compute, nil
}