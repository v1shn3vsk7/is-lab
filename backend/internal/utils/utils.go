package utils

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/v1shn3vsk7/is-lab/internal/tech/closer"
)

const gracefulShutdownWaitTime = 2 * time.Second

func GracefulShutdown(cancel context.CancelFunc) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(ch)

	sig := <-ch
	errorMessage := fmt.Sprintf("%s %v - %s", "Received shutdown signal:", sig, "Graceful shutdown done")
	log.Info().Msg(errorMessage)
	cancel()

	closer.CloseAll()

	time.Sleep(gracefulShutdownWaitTime)
}
