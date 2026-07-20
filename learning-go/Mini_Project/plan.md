# Plan: Fleet Manager demo implementation

## Problem statement
Implement the Fleet Manager demo described in learning-go/Mini_Project/README.md. Deliver a working CLI demo (cmd/fleet) that demonstrates domain modeling (drones, packages, locations), an in-memory store, fleet manager logic, custom errors, and a simple scenario: register drones, create packages, load/move/unload, and print final state.

## Proposed approach
1. Implement domain models in pkg/models: Entity, Location, Package, Drone. Add methods on Drone (Load, Unload, MoveTo) with pointer receivers and a value-receiver Summary.
2. Implement custom errors in pkg/errors (FleetError type + sentinel errors).
3. Implement an in-memory store in pkg/store that satisfies the Storage interface (maps, simple helpers).
4. Implement Fleet manager in pkg/fleet: NewFleet, RegisterDrone, AssignPackageToBestDrone, helper functions for selecting drones.
5. Implement CLI in cmd/fleet/main.go to create store & fleet, add drones & packages, run demo scenario, and print readable outputs and errors.
6. Add basic unit tests for models, store, and fleet manager.

## Deliverables
- pkg/models: entity.go, location.go, package.go, drone.go (implemented)
- pkg/errors: errors.go (FleetError + sentinel errors)
- pkg/store: memory.go (in-memory implementation)
- pkg/fleet: manager.go (fully implemented using models and store)
- cmd/fleet/main.go: demo program
- README update or CLI usage printed on run

## Todos
- plan-implement-models: Implement domain models (Entity, Location, Package, Drone) and Drone methods.
- plan-implement-errors: Implement pkg/errors (FleetError type and sentinel errors).
- plan-implement-store: Implement in-memory store (pkg/store/memory.go) satisfying Storage.
- plan-implement-fleet-manager: Implement Fleet methods and selection logic.
- plan-implement-cli: Implement cmd/fleet demo demonstrating the scenario.
- plan-add-tests: Add unit tests for models, store, and fleet manager.

## Notes & considerations
- Keep implementation single-threaded and use only the stdlib.
- Use pointer receivers for mutating methods; value receiver for summary/display methods.
- Battery consumption: simple formula (euclidean distance * 1% per unit + flat 5%). Document in code.
- ID generation: use simple fmt-based ids (e.g., drone-1, pkg-1) to avoid external deps.
- Ensure errors are typed so callers can assert on sentinel errors.

## Verification
- Build: `go build ./...`
- Run demo: `go run ./cmd/fleet`
- Expected output: operations logged, errors printed where expected, final state showing battery and package status changes.

## Milestones
1. Models + errors (complete)  
2. Store + fleet manager (complete)  
3. CLI demo + tests (complete)

---
Created-by: Copilot plan generator
