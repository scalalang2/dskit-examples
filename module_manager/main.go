package main

import (
	"context"
	"os"

	"github.com/go-kit/log"
	"github.com/grafana/dskit/modules"
	"github.com/grafana/dskit/services"
)

var logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))

type InitService func() (services.Service, error)

func initServiceA(name string) InitService {
	return func() (services.Service, error) {
		logger.Log("msg", "service initialized", "name", name)
		return services.NewIdleService(func(serviceContext context.Context) error {
			logger.Log("msg", "starting service", "name", name)
			return nil
		}, func(failureCase error) error {
			logger.Log("msg", "stopping service", "name", name)
			return nil
		}), nil
	}
}

func main() {
	mm := modules.NewManager(log.NewNopLogger())
	mm.RegisterModule("serviceA", initServiceA("serviceA"))
	mm.RegisterModule("serviceB", initServiceA("serviceB"))
	mm.RegisterModule("serviceC", initServiceA("serviceC"))
	mm.RegisterModule("serviceD", initServiceA("serviceD"))
	deps := map[string][]string{
		"serviceA": {"serviceB"},
		"serviceB": {"serviceC"},
	}
	for mod, targets := range deps {
		if err := mm.AddDependency(mod, targets...); err != nil {
			panic(err)
		}
	}

	svc, err := mm.InitModuleServices("serviceA")
	if err != nil {
		panic(err)
	}
	for name, val := range svc {
		logger.Log("state", val.State(), "name", name)
	}

	svcs := []services.Service(nil)
	for _, s := range svc {
		svcs = append(svcs, s)
	}
	serviceManager, err := services.NewManager(svcs...)
	if err != nil {
		panic(err)
	}

	err = services.StartManagerAndAwaitHealthy(context.Background(), serviceManager)
	if err != nil {
		panic(err)
	}

	for name, val := range svc {
		logger.Log("state", val.State(), "name", name)
	}

	err = services.StopManagerAndAwaitStopped(context.Background(), serviceManager)
	if err != nil {
		panic(err)
	}
	logger.Log("msg", "all services terminated")
}
