package keeper

import (
	"errors"
	"fmt"
	"time"
)

var (
	KEEPER_ERR_RESOURCE_ALREADY_EXISTS = errors.New("resource already exists in pool")
	KEEPER_ERR_ALREADY_RUNNING         = errors.New("keeper already started")
)

type Keeper struct {
	// Interval at which resources should be checked.
	interval time.Duration
	// retries after failed respawning
	respawnRetry uint8
	// retries after failed healtcheck
	failRetry uint8
	// The retry strategy, plain, exponential backoff, etc.
	backoff Backoff
	// the pool of resources
	resourcePool map[string]Resource
	// done channel to indicate stop
	done chan struct{}
}

func (k *Keeper) String() string {
	return fmt.Sprintf("Keeper [interval:%d, respawnRetry:%d, failRetry:%d, backoff:%v ]",
		k.interval, k.respawnRetry, k.failRetry, k.backoff)
}

type KeeperOpt func(*Keeper)

// NewKeeper returns a new keeper instance with supplied options
func NewKeeper(keeperOpts ...KeeperOpt) *Keeper {
	keeper := &Keeper{}
	for _, opt := range keeperOpts {
		opt(keeper)
	}
	return keeper
}

// WithInterval time to wait between retries
func WithInterval(interval time.Duration) KeeperOpt {
	return func(k *Keeper) {
		k.interval = interval
	}
}

// WithFailRetry how many times to retry when health check fails
func WithFailRetry(failRetry uint8) KeeperOpt {
	return func(k *Keeper) {
		k.failRetry = failRetry
	}
}

// WithRespawnRetry how many times to retry when respawn fails
func WithRespawnRetry(respawnRetry uint8) KeeperOpt {
	return func(k *Keeper) {
		k.respawnRetry = respawnRetry
	}
}

// WithBackoff backoff strategy to apply during retry / healtcheck failure
func WithBackoff(backoff Backoff) KeeperOpt {
	return func(k *Keeper) {
		k.backoff = backoff
	}
}

func (k *Keeper) Start() error {
	if k.done != nil {
		return KEEPER_ERR_ALREADY_RUNNING
	}
	k.done = make(chan struct{})
	k.start()
	return nil
}

// Stop stops process and health check loop
func (k *Keeper) Stop() {
	k.done <- struct{}{}
}

func (k *Keeper) start() {
	go func() {
		ticker := time.NewTicker(k.interval)
		defer ticker.Stop()
		for {
			select {
			case <-k.done:
				fmt.Println("received process stop")
				// stop processing
				return
			case <-ticker.C:
				fmt.Println("received tikcer checking health")
				// start healthcheck for items in pool
				go k.check()
			}
		}
	}()
}

func (k *Keeper) check() {
	for name, resource := range k.resourcePool {
		fmt.Printf("calling health check on %s\n", name)
		err := resource.Check()
		if err != nil {
			fmt.Printf("check resulted in error for %s with %v\n", name, err)
		}
	}
}
func (k *Keeper) AddResource(r Resource) error {
	if k.resourcePool == nil {
		k.resourcePool = make(map[string]Resource)
	}
	if _, ok := k.resourcePool[r.Name()]; ok {
		return KEEPER_ERR_RESOURCE_ALREADY_EXISTS
	}
	rsc, _ := r.Spawn()
	k.resourcePool[r.Name()] = rsc
	return nil
}
