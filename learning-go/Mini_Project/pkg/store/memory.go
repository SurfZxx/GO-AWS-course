package store

// MemoryStore is a simple in-memory implementation of the Storage interface.
// Replace interface{} with concrete types from pkg/models when implementing.

type MemoryStore struct {
	Drones   map[string]interface{}
	Packages map[string]interface{}
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{Drones: make(map[string]interface{}), Packages: make(map[string]interface{})}
}

func (m *MemoryStore) SaveDrone(d interface{}) error { m.Drones["TODO"] = d; return nil }
func (m *MemoryStore) GetDrone(id string) (interface{}, error) { return m.Drones[id], nil }
func (m *MemoryStore) SavePackage(p interface{}) error { m.Packages["TODO"] = p; return nil }
func (m *MemoryStore) GetPackage(id string) (interface{}, error) { return m.Packages[id], nil }
