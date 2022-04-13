package models

import (
	"time"
)

// User struct to describe User object.
type GCEInstance struct {
	ID           uint64     `db:"id" json:"id" validate:"required"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
	Name         string    `db:"name" json:"name" validate:"required"`
	Zone         string    `db:"zone" json:"zone" validate:"required"`
	Status       string    `db:"status" json:"status" validate:"required"`
	Compute      Compute   `db:"compute_id" json:"computeId"`
}
