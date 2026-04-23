package models

import "time"

// Entity is an embeddable base for domain objects.
type Entity struct {
	ID        string
	CreatedAt time.Time
}
