package signals

import (
	"os"
	"os/signal"
)

var onlyOneSignalHandler = make(chan struct{})

func SetupSignalHandler() (stopCh <-chan struct{}) {
	// panics when called twice
	close(onlyOneSignalHandler)

	stop := make(chan struct{})
	signalCh := make(chan os.Signal, 2)
	signal.Notify(signalCh, shutdownSignals...)
	go func() {
		<-signalCh
		close(stop)
		<-signalCh
		// second signal, exit directly.
		os.Exit(1)
	}()

	return stop
}
