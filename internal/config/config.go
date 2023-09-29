package config

import (
	"os"
	"time"

	"github.com/Ahmed-Sermani/driver-tracking-system/internal/rpc"
	"github.com/Ahmed-Sermani/driver-tracking-system/internal/web"
	"github.com/kelseyhightower/envconfig"

	"github.com/stackus/dotenv"
)

type (
	PGConfig struct {
		Conn string `required:"true"`
	}

	JWT struct {
		Secret string        `required:"true"`
		Expiry time.Duration `required:"true" default:"24h"`
	}

	Cent struct {
		Conn string `required:"true"`
		Key  string `required:"true"`
	}

	AppConfig struct {
		Environment     string
		LogLevel        string `envconfig:"LOG_LEVEL" default:"DEBUG"`
		PG              PGConfig
		Cent            Cent
		JWT             JWT
		Rpc             rpc.RpcConfig
		Web             web.WebConfig
		ShutdownTimeout time.Duration `envconfig:"SHUTDOWN_TIMEOUT" default:"30s"`
	}
)

func InitConfig() (cfg AppConfig, err error) {
	if err = dotenv.Load(dotenv.EnvironmentFiles(os.Getenv("ENVIRONMENT"))); err != nil {
		return
	}

	err = envconfig.Process("", &cfg)

	return
}
