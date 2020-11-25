package util

import (
	"context"
	"os"
	"os/signal"
)

func WithSignal(parent context.Context, sig ...os.Signal) (ctx context.Context, cancel context.CancelFunc) {
	ctx, cancel = context.WithCancel(parent)
	c := make(chan os.Signal, 1)
	signal.Notify(c, sig...)
	go func() {
		sig := <-c
		Logf("Caught %v. Send signal again to interupt immediately. Canceling...", sig)
		signal.Stop(c)
		cancel()
	}()
	return
}
