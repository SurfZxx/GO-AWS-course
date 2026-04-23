package fleet

// Fleet contains drones and a storage backend. This file contains a minimal type sketch.

// Note: models.* types are declared in pkg/models; replace the placeholder types/imports when implementing.

type Storage interface {
	SaveDrone(d interface{}) error
	GetDrone(id string) (interface{}, error)
	SavePackage(p interface{}) error
	GetPackage(id string) (interface{}, error)
}

type Fleet struct {
	Store  Storage
	Drones map[string]interface{} // replace interface{} with *models.Drone
}

func NewFleet(store Storage) *Fleet {
	return &Fleet{Store: store, Drones: make(map[string]interface{})}
}

func (f *Fleet) RegisterDrone(d interface{}) error { return nil }
func (f *Fleet) AssignPackageToBestDrone(p interface{}) (string, error) { return "", nil }
