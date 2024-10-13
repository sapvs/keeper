package sample

import (
	"errors"
	"fmt"

	"github.com/sapvs/keeper"
)

type DummyResource struct {
	checkCount  uint8
	failOnCheck uint8
}

// Health check, whether remake is needed
func (c *DummyResource) Check() error {
	c.checkCount++
	if c.checkCount == c.failOnCheck {
		c.checkCount = 0
		return errors.New("failing intentionally")
	}
	return nil
}

// Name of the resource with which it will be stored and retrieved
func (c *DummyResource) Name() string {
	return "DummyResource"
}

// Spawn creates/rec-creates resource, e.g. reconnect to database.
func (*DummyResource) Spawn() (keeper.Resource, error) {
	return &DummyResource{failOnCheck: 3}, nil
}

// Error func to handle error in Check or ReMake
func (c *DummyResource) Error(err keeper.Keeper_Error, message any) {
	fmt.Printf("failed resource with %d due to %s\n", err, message)
}

var _ keeper.Resource = &DummyResource{}
