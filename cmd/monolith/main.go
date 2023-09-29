package main

import (
	"database/sql"
	"fmt"
	"io/fs"
	"net/http"
	"os"

	"github.com/Ahmed-Sermani/driver-tracking-system/emulation"
	"github.com/Ahmed-Sermani/driver-tracking-system/internal/auth"
	"github.com/Ahmed-Sermani/driver-tracking-system/internal/config"
	"github.com/Ahmed-Sermani/driver-tracking-system/internal/logger"
	"github.com/Ahmed-Sermani/driver-tracking-system/internal/monolith"
	"github.com/Ahmed-Sermani/driver-tracking-system/internal/waiter"
	"github.com/Ahmed-Sermani/driver-tracking-system/internal/web"
	"github.com/Ahmed-Sermani/driver-tracking-system/proxy"
	"github.com/Ahmed-Sermani/driver-tracking-system/users"
	usersMigrations "github.com/Ahmed-Sermani/driver-tracking-system/users/migrations"
	"github.com/go-chi/chi/v5"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/pressly/goose/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func run() (err error) {
	var cfg config.AppConfig
	cfg, err = config.InitConfig()
	if err != nil {
		return err
	}

	m := app{cfg: cfg}

	// init infrastructure...
	m.db, err = sql.Open("pgx", cfg.PG.Conn)
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			return
		}
	}(m.db)

	// migrations
	migrations := []fs.FS{
		usersMigrations.FS,
	}
	for i := range migrations {
		if err := migrateDB(m.db, migrations[i]); err != nil {
			return err
		}
	}

	m.logger = logger.New(logger.LogConfig{
		Environment: cfg.Environment,
		LogLevel:    logger.Level(cfg.LogLevel),
	})
	m.rpc = initRpc(cfg)
	m.mux = initMux(cfg.Web)
	m.waiter = waiter.New(waiter.CatchSignals())

	// init modules
	m.modules = []monolith.Module{
		new(proxy.Module),
		new(users.Module),
		new(emulation.Module),
	}

	if err = m.startupModules(); err != nil {
		return err
	}

	// Mount general web resources
	m.mux.Mount("/", http.FileServer(http.FS(web.WebUI)))

	// health check
	m.mux.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	})

	fmt.Println("started application")
	defer fmt.Println("stopped application")

	m.waiter.Add(
		m.waitForWeb,
		m.waitForRPC,
	)

	return m.waiter.Wait()
}

func initRpc(cfg config.AppConfig) *grpc.Server {
	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			auth.NewUnaryAuthInterceptor(
				cfg.JWT.Secret,
				[]string{"/userspb.UsersService/Login", "/userspb.UsersService/RegisterUser"},
				[]string{"id"},
			),
		),
		grpc.ChainStreamInterceptor(
			auth.NewStreamAuthInterceptor(
				cfg.JWT.Secret,
				[]string{},
				[]string{"id"},
			),
		),
	)
	reflection.Register(server)

	return server
}

func initMux(_ web.WebConfig) *chi.Mux {
	return chi.NewMux()
}

func migrateDB(db *sql.DB, fs fs.FS) error {
	goose.SetBaseFS(fs)
	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}
	goose.Status(db, ".")
	return goose.Up(db, ".")
}
