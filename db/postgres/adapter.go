package postgres

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Config struct {
	Host     string `koanf:"host"`
	Username string `koanf:"username"`
	Password string `koanf:"password"`
	Name     string `koanf:"name"`
}

type Postgres struct {
	db  *gorm.DB
	cfg Config
}

func NewPostgresAdapter(cfg Config) PgAdapter {
	return Postgres{db: &gorm.DB{}, cfg: cfg}
}

func (p Postgres) getInstance(uri string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(uri))
	if err != nil {
		log.Fatalf(err.Error())
	}

	return db
}

func (p Postgres) ConnectInstance() *gorm.DB {
	// TODO - config all these database connecting string
	dsn := fmt.Sprintf("postgres://%s:%s@%s/%s", p.cfg.Username, p.cfg.Password, p.cfg.Host, p.cfg.Name)
	db := p.getInstance(dsn)

	return db
}

func (p Postgres) retryHandler(int, func() (bool, error)) error {
	return nil
}

func (p Postgres) setConnectionPool(d *gorm.DB) {
	db, err := d.DB()
	if err != nil {
		panic(err)
	}
	// TODO - config these
	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(20)
	db.SetConnMaxIdleTime(20)
}
