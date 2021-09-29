package database

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type BaseDatabaseInterface interface {
	Setup()
	GetDB() *gorm.DB
}
type DatabaseInterface interface {
	BaseDatabaseInterface
	// RunMigration()
	InitialBootstrap()
}

var db *gorm.DB

func NewPostgresDatabase() DatabaseInterface {
	return &Database{}
}

type Database struct{}

func (d *Database) Setup() {
	dsn := os.Getenv("POSTGRES_URI")
	fmt.Println("dsn", dsn)
	var newDB *gorm.DB
	var err error

	newDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	var retries = 10
	for err != nil {
		if retries > 1 {
			retries--
			time.Sleep(5 * time.Second)
			newDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
			continue
		}
		panic(err)
	}

	db = newDB

	d.InitialBootstrap()

	// d.RunMigration()
}

func (*Database) GetDB() *gorm.DB {
	return db
}

func (d *Database) InitialBootstrap() {
	// add the uuid extension
	db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)

	// add the account schema
	db.Exec("CREATE SCHEMA IF NOT EXISTS account AUTHORIZATION admin; GRANT ALL ON SCHEMA account TO PUBLIC; GRANT ALL ON SCHEMA public TO admin;")
	db.Exec("CREATE SCHEMA IF NOT EXISTS engagement AUTHORIZATION admin; GRANT ALL ON SCHEMA engagement TO PUBLIC; GRANT ALL ON SCHEMA public TO admin;")
}
