package keeper

import (
	"testing"
	"time"
)

func TestNewKeeper(t *testing.T) {
	type args struct {
		interval     time.Duration
		respawnRetry uint8
		failRetry    uint8
	}
	tests := []struct {
		name string
		args args
	}{
		{"Keeper-1", args{interval: 10 * time.Second, respawnRetry: 23, failRetry: 21}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			kper := NewKeeper(WithFailRetry(tt.args.failRetry),
				WithInterval(tt.args.interval),
				WithRespawnRetry(tt.args.respawnRetry))

			if kper.interval != tt.args.interval ||
				kper.respawnRetry != tt.args.respawnRetry ||
				kper.failRetry != tt.args.failRetry {
				t.Errorf("Failed %s\n", kper)
			}
		})
	}
}
