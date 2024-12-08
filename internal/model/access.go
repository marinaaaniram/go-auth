package model

import (
	"time"
)

// Internal Access model
type Access struct {
	ID             int64
	EndpointAdress string
	CreatedAt      time.Time
	UpdatedAt      *time.Time
}
