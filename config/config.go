package config

import (
	"github.com/doo-dev/pech-pech/db/postgres"
	"github.com/doo-dev/pech-pech/infrastructure/mail"
	authConfig "github.com/doo-dev/pech-pech/internal/modules/auth/usecase"
	"github.com/doo-dev/pech-pech/internal/server/api"
)

type Config struct {
	Debug      bool              `koanf:"debug"`
	AuthConfig authConfig.Config `koanf:"auth_config"`
	PgDB       postgres.Config   `koanf:"pg_db"`
	MailConfig mail.Config       `koanf:"mail_config"`
	HttpServer api.Config        `koanf:"http_server"`
}
