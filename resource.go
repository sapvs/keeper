package keeper

// Resource implementations of this interface are monitored
type Resource interface {
	// Name of the resource with which it will be stored and retrieved
	Name() string
	// Health check, whether remake is needed
	Check() error
	// Spawn creates/rec-creates resource, e.g. reconnect to database.
	Spawn() (Resource, error)
	//Error func to handle error in Check or ReMake
	Error(Keeper_Error, any)
}
