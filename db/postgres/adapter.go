package postgres

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Postgres struct {
	db *gorm.DB
}

func NewPostgresAdapter() PgAdapter {
	return Postgres{db: &gorm.DB{}}
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
	dsn := fmt.Sprintf("postgres://%s:%s@%s?sslmod", "doo-dev", "123456", "localhost")
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
