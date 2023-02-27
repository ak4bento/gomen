package database

import (
	"database/sql"
	"fmt"
	"gomen/app/exceptions"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

// Config for databases config
type Config struct {
	//DATABASE
	HOST string `env:"DB_HOST" envDefault:"localhost" json:"db_host"`
	PORT string `env:"DB_PORT" envDefault:"27107" json:"db_port"`
	NAME string `env:"DB_NAME" json:"db_name"`
	USER string `env:"DB_USER" json:"db_user"`
	PASS string `env:"DB_PASS" json:"db_pass"`
	SSL  string `env:"DB_SSL" envDefault:"disable" json:"db_ssl"`
	TZ   string `env:"TIME_ZONE" envDefault:"Asia/Jakarta" json:"db_tz"`

	MAX_IDLE_CONNS      int `env:"MAX_IDLE_CONNS" json:"max_idle_conns"`           //default 100
	MAX_OPEN_CONNS      int `env:"MAX_OPEN_CONNS" json:"max_open_conns"`           //default 1000
	MAX_IDLE_CONNS_TIME int `env:"MAX_IDLE_CONNS_TIME" json:"max_idle_conns_time"` //default 60m

	//ETC
	EngineName string `env:"ENGINE_NAME" json:"engine_name,omitempty"`
}

type PostgreSQL struct {
	Config Config
}

func (pgsql *PostgreSQL) Connect(cfg *gorm.Config) (DB *gorm.DB, sql_ *sql.DB) {

	var (
		dsn string
		err error
	)

	//Set DSN connection
	// dbname for postgres version <= 14 and database for postgres above 14
	dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s database=%s port=%s sslmode=%s TimeZone=%s",
		pgsql.Config.HOST,
		pgsql.Config.USER,
		pgsql.Config.PASS,
		pgsql.Config.NAME, //for postgres version <= 14
		pgsql.Config.NAME, //for postgres above 14
		pgsql.Config.PORT,
		pgsql.Config.SSL,
		pgsql.Config.TZ,
	)

	//set gorm config if exist
	if cfg == nil {
		cfg = &gorm.Config{}
	}

	//create connection
	if DB, err = gorm.Open(postgres.Open(dsn), cfg); err != nil {
		exceptions.BadRequestException(err)
		panic(err)
	}

	//get database
	if sql_, err = DB.DB(); err != nil {
		exceptions.BadRequestException(err)
		panic(err)
	}

	if pgsql.Config.MAX_IDLE_CONNS_TIME > 0 {
		sql_.SetConnMaxIdleTime(time.Duration(pgsql.Config.MAX_IDLE_CONNS_TIME) * time.Minute)
	}
	if pgsql.Config.MAX_IDLE_CONNS_TIME > 0 {
		sql_.SetMaxIdleConns(pgsql.Config.MAX_IDLE_CONNS)
	}
	if pgsql.Config.MAX_OPEN_CONNS > 0 {
		sql_.SetMaxOpenConns(pgsql.Config.MAX_OPEN_CONNS)
	}

	return
}
