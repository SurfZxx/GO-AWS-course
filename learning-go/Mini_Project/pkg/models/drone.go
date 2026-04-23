package models

// Drone models a delivery drone and implements Carrier (pointer receiver methods).
type Drone struct {
	Entity
	Battery    int // 0..100
	Location   Location
	CapacityKg float64
	Cargo      []*Package
	Status     string // e.g., "idle", "flying", "charging"
}

// Method sketches (implementations omitted / minimal for the sketch).
func (d *Drone) Load(p *Package) error { return nil }
func (d *Drone) Unload(pkgID string) (*Package, error) { return nil, nil }
func (d *Drone) MoveTo(dest Location) error { return nil }
func (d Drone) Summary() string { return "" }
