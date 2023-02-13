package app

import (
	"context"
	"dskit-examples/utils"
	"fmt"
	"net/http"
	"os"

	"github.com/go-kit/log"
	"github.com/grafana/dskit/kv"
	"github.com/grafana/dskit/modules"
	"github.com/grafana/dskit/ring"
	"github.com/grafana/dskit/services"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/weaveworks/common/server"
)

var Logger = log.With(
	log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr)),
	"ts", log.DefaultTimestampUTC,
	"caller", log.DefaultCaller,
)

type App struct {
	registry *prometheus.Registry
	srv      *server.Server
	ring     *ring.Ring

	ModuleManager *modules.Manager
}

func NewApp() (*App, error) {
	app := &App{}
	mm := modules.NewManager(Logger)
	mm.RegisterModule("server", app.initServer, modules.UserInvisibleModule)
	mm.RegisterModule("ring", app.initRing, modules.UserInvisibleModule)

	if err := mm.AddDependency("ring", "server"); err != nil {
		return nil, fmt.Errorf("failed to add dependency: %w", err)
	}

	app.ModuleManager = mm

	registry := prometheus.NewRegistry()
	app.registry = registry

	return app, nil
}

func (a *App) Run() error {
	serviceMap, err := a.ModuleManager.InitModuleServices("ring", "server")
	if err != nil {
		return err
	}

	Logger.Log("msg", "all modules are initialized")

	var servs []services.Service
	for _, s := range serviceMap {
		servs = append(servs, s)
	}

	sm, err := services.NewManager(servs...)
	if err != nil {
		return err
	}

	err = sm.StartAsync(context.Background())
	if err != nil {
		return err
	}

	return sm.AwaitStopped(context.Background())
}

func (a *App) initServer() (services.Service, error) {
	cfg := server.Config{
		HTTPListenNetwork: server.NetworkTCPV4,
		HTTPListenAddress: "localhost",
		HTTPListenPort:    9290,
	}

	srv, err := server.New(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create server: %w", err)
	}

	servicesToWaitFor := func() []services.Service {
		return nil
	}

	a.srv = srv
	a.srv.HTTP.Handle("/hello", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	}))

	return utils.NewServerService(srv, servicesToWaitFor, Logger), nil
}

func (a *App) initRing() (services.Service, error) {
	kvConfig := kv.Config{Store: "inmemory", Prefix: "collectors/"}
	cfg := ring.Config{
		KVStore:              kvConfig,
		HeartbeatTimeout:     0,
		ReplicationFactor:    3,
		ZoneAwarenessEnabled: true,
	}

	store, err := kv.NewClient(kvConfig, ring.GetCodec(), a.registry, Logger)
	if err != nil {
		return nil, err
	}

	r, err := ring.NewWithStoreClientAndStrategy(cfg, "test-name", "test-key", store, ring.NewDefaultReplicationStrategy(), a.registry, Logger)
	if err != nil {
		return nil, err
	}
	a.srv.HTTP.Handle("/ring", r)
	a.ring = r
	return a.ring, nil
}
