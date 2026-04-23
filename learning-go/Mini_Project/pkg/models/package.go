package models

// Package represents a delivery package in the domain.
// Field names follow the README's suggested sketch.
type Package struct {
	Entity
	Recipient   string
	WeightKg    float64
	Destination Location
	Status      string // e.g., "waiting", "in_transit", "delivered"
}
