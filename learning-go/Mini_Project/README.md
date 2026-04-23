# Fleet Manager — Mini Project (Go OOP Practice)

**Difficulty: Moderate**

---

## Learning Objectives
- Define and initialize structs
- Add methods with value and pointer receivers; understand when to use each
- Design and implement interfaces
- Use composition and embedding to share behavior and fields
- Organize code into packages
- Create and use custom error types and handle errors idiomatically

---

## Project Title
Fleet Manager — Delivery Drone Fleet Simulator

## Problem Statement
Build a small fleet management library and CLI demo for delivery drones. Each drone carries packages between locations, consumes battery, and reports status. The goal is to model domain types using Go structs, behaviors with methods and interfaces, and package boundaries for clean design. Include clear error handling using custom error types.

This is not a real-time simulator — focus on modeling, OOP patterns in Go, and predictable behavior.

---

## High-level Requirements

### 1. Design Domain Types:
- **Entity**: common embedded type with ID, CreatedAt
- **Location**: latitude & longitude (float64), Name (string)
- **Package**: recipient name, weight (kg), destination Location, status
- **Drone**: ID, battery level (0–100), location, capacity (kg), cargo (slice of *Package), status

### 2. Interfaces:
- **Carrier** interface with methods:
  - `Load(pkg *Package) error`
  - `Unload(pkgID string) (*Package, error)`
  - `MoveTo(dest Location) error`
  - `Status() string`
- **Storage** interface (for packages and drones persistence) with methods:
  - `SaveDrone(d *Drone) error`
  - `GetDrone(id string) (*Drone, error)`
  - `SavePackage(p *Package) error`
  - `GetPackage(id string) (*Package, error)`

### 3. Methods:
- Implement Carrier on Drone using pointer receiver where methods mutate drone state (Load, Unload, MoveTo).
- Implement at least one value-receiver method on Drone (e.g., `Summary() string`) to demonstrate non-mutating behavior.
- Implement Stringer-like methods that return readable descriptions.

### 4. Composition / Embedding:
- Embed Entity in Drone and Package.
- Use composition for Fleet: a struct that holds a list of drones and a Storage implementation.

### 5. Package Organization:
- `cmd/fleet` — small CLI tool demonstrating functionality (e.g., create fleet, add drones/packages, simulate a trip)
- `pkg/models` — domain types (Entity, Drone, Package, Location)
- `pkg/fleet` — Fleet management logic and Carrier implementations
- `pkg/store` — an in-memory storage implementation that satisfies Storage
- `pkg/errors` — custom error types and exported sentinel errors

### 6. Error Handling:
- Implement domain errors as custom types (e.g., `type FleetError struct {...}`) that implement error.
- Provide sentinel errors for common conditions (ErrNotFound, ErrOverCapacity, ErrLowBattery, ErrAlreadyLoaded).
- Methods must return errors when operations fail; caller should check and handle them.

### 7. Deliver a Simple Demo (cmd/fleet/main.go) that:
- Creates a fleet and storage
- Adds two drones
- Creates two packages with destinations
- Loads package onto a drone, attempts to move it; show error handling if battery insufficient
- Prints final states (drones and packages)

---

## Constraints
- Keep implementation single-threaded (no goroutines required)
- Use only Go standard library
- Use Go modules (init a module root)
- Keep code small and idiomatic
- Do not persist to external DB (in-memory store required; file persistence is a bonus)
- Methods that mutate state must use pointer receivers

---

## Success Criteria
- All domain types implemented and organized into packages as specified
- Drone implements Carrier interface
- Storage interface implemented by an in-memory store (pkg/store)
- Custom error types implemented and used in relevant scenarios
- CLI demo builds and runs, demonstrating:
  - Loading a package
  - Handling OverCapacity or LowBattery error
  - Moving a drone and updating location and battery
  - Unloading a package and updating package status
- Code compiles: `go build ./...`
- README (or CLI usage printed) explains how to run the demo

Optional verification you can run:
- `go run ./cmd/fleet`
- Output should clearly show operations, and printed errors where expected
- Final printed state should reflect successful operations (battery decreased, package status changed)

---

## Suggested File Layout (Starter)
```
.
├── go.mod
├── README.md
├── cmd/
│   └── fleet/
│       └── main.go
└── pkg/
    ├── models/
    │   ├── entity.go
    │   ├── location.go
    │   ├── package.go
    │   └── drone.go
    ├── fleet/
    │   └── manager.go
    ├── store/
    │   └── memory.go
    └── errors/
        └── errors.go
```

---

## API / Type Sketch & Example Signatures

### pkg/models/entity.go
```go
type Entity struct {
    ID        string
    CreatedAt time.Time
}
```

### pkg/models/location.go
```go
type Location struct {
    Name string
    Lat  float64
    Lon  float64
}
```

### pkg/models/package.go
```go
type Package struct {
    Entity
    Recipient   string
    WeightKg    float64
    Destination Location
    Status      string // e.g., "waiting", "in_transit", "delivered"
}
```

### pkg/models/drone.go
```go
type Drone struct {
    Entity
    Battery    int          // 0..100
    Location   Location
    CapacityKg float64
    Cargo      []*Package
    Status     string       // e.g., "idle", "flying", "charging"
}

func (d *Drone) Load(p *Package) error
func (d *Drone) Unload(pkgID string) (*Package, error)
func (d *Drone) MoveTo(dest Location) error
func (d Drone) Summary() string // value receiver example
```

### pkg/fleet/manager.go
```go
type Fleet struct {
    Store  Storage
    Drones map[string]*Drone
}

func NewFleet(store Storage) *Fleet
func (f *Fleet) RegisterDrone(d *Drone) error
func (f *Fleet) AssignPackageToBestDrone(p *Package) (string, error)
```

### pkg/store/memory.go
```go
type MemoryStore struct {
    drones   map[string]*Drone
    packages map[string]*Package
}
// Implement Save/Get methods
```

### pkg/errors/errors.go
```go
type FleetError struct {
    Code    int
    Message string
}

func (e FleetError) Error() string

var (
    ErrNotFound      = FleetError{Code: 404, Message: "not found"}
    ErrOverCapacity  = FleetError{Code: 422, Message: "over capacity"}
    ErrLowBattery    = FleetError{Code: 423, Message: "insufficient battery"}
    ErrAlreadyLoaded = FleetError{Code: 409, Message: "already loaded"}
)
```

---

## Hints
- Use pointer receivers for methods that change the drone (Load, Unload, MoveTo). Use value receivers for read-only display-like methods.
- When checking capacity, sum current cargo weight + new package weight.
- Moving reduces battery: define a simple battery consumption formula (e.g., distance * 1% per km + fixed 5%).
- For distances, you can compute Euclidean approximation (no need for haversine).
- For unique IDs, use simple unique strings (e.g., `fmt.Sprintf("drone-%d", n)` or use `github.com/google/uuid` if you choose—but stdlib is fine).
- Make custom errors comparable by type or assert to `*FleetError` for richer checks.
- Keep Storage interfaces small and simple; the in-memory store can use maps.

---

## Bonus Challenges (Optional)
1. Persist store to JSON files (file-based store package).
2. Add command-line flags to the demo to create random drones and packages.
3. Implement scheduling: Fleet automatically picks the nearest available drone with sufficient battery and capacity to deliver a package.
4. Add unit tests for models, store, and fleet manager.
5. Add concurrency: allow multiple assignments and demonstrate safe access to MemoryStore (use mutexes).

---

## Example Demo Scenario (What the CLI Should Show)
- Create fleet with MemoryStore
- Register Drone A (capacity 5kg, battery 80) and Drone B (capacity 2kg, battery 40)
- Create Package 1 (weight 3kg, destination loc X) and Package 2 (weight 2.5kg, destination loc Y)
- Assign Package 1 to Drone A (succeeds)
- Try to assign Package 2 to Drone B (fails with ErrOverCapacity) — show error handling
- Move Drone A to destination (battery reduced). If battery would be insufficient, MoveTo returns ErrLowBattery.
- Unload package at destination; package status becomes "delivered"
- Print final state

---

## Good Luck!
This project will test your understanding of Go's OOP patterns and package organization. Focus on clean design, proper error handling, and using the right receiver type (value vs. pointer) for each method. Once you complete this, you'll have a solid grasp of building maintainable Go applications!
