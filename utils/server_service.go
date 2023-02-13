package utils

import (
	"context"
	"fmt"

	"github.com/go-kit/log"
	"github.com/grafana/dskit/services"
	"github.com/weaveworks/common/server"
)

func NewServerService(serv *server.Server, servicesToWaitFor func() []services.Service, logger log.Logger) services.Service {
	serverDone := make(chan error, 1)

	runFn := func(ctx context.Context) error {
		go func() {
			defer close(serverDone)
			serverDone <- serv.Run()
		}()

		logger.Log("msg", "server is ready to handle requests", "address", serv.HTTPServer.Addr)

		select {
		case <-ctx.Done():
			return nil
		case err := <-serverDone:
			if err != nil {
				return err
			}
			return fmt.Errorf("server stopped unexpectedly")
		}
	}

	stoppingFn := func(_ error) error {
		for _, s := range servicesToWaitFor() {
			_ = s.AwaitTerminated(context.Background())
		}

		logger.Log("msg", "server is terminated")

		// shutdown HTTP and gRPC servers (this also unblocks Run)
		serv.Shutdown()

		<-serverDone
		return nil
	}

	return services.NewBasicService(nil, runFn, stoppingFn)
}
