package keeper

type KeeperError error

type Keeper_Error int

const (
	//ALL_TRIES_FAILED failed to maintain Resource after configured retries. Critical Error
	ALL_TRIES_FAILED Keeper_Error = iota + 1
	// CREATE_FAILED Failed to create, but keeper will still retry
	CREATE_FAILED
	// CHECK_FAILED helathcheck has failed, keeper will retry
	CHECK_FAILED
)


