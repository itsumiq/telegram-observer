package user

// UUIDer represent contract that user UUIDers should implement.
type UUIDer interface {
	// Create creates uuid and returns it.
	Create() string
}
