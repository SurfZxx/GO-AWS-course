package errors

import "fmt"

// FleetError is a custom error type for domain errors.
type FleetError struct {
	Code    int
	Message string
}

func (e FleetError) Error() string { return fmt.Sprintf("%d: %s", e.Code, e.Message) }

var (
	ErrNotFound      = FleetError{Code: 404, Message: "not found"}
	ErrOverCapacity  = FleetError{Code: 422, Message: "over capacity"}
	ErrLowBattery    = FleetError{Code: 423, Message: "insufficient battery"}
	ErrAlreadyLoaded = FleetError{Code: 409, Message: "already loaded"}
)
