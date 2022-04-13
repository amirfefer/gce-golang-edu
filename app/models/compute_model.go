package models

import (
	"time"
)

// Compute struct to describe compute object.
type Compute struct {
	ID           int       `db:"id" json:"id" validate:"required"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
	Name         string    `db:"name" json:"name" validate:"required"`
	JsonPath     string    `db:"json_path" json:"path" validate:"required"`
	ProjectID    string    `db:"project_id" json:"projectID" validate:"required"`
	Zone         string    `db:"zone" json:"zone" validate:"required"`
}
