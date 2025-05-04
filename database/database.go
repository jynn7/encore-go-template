package database

import (
	"encore.dev/storage/sqldb"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//encore:service
type Service struct {
	db *gorm.DB
}

// you can name your database whatever you want
var myDB = sqldb.NewDatabase("my_db_name", sqldb.DatabaseConfig{
	Migrations: "./migrations",
})

func InitService() (*Service, error) {

	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: myDB.Stdlib(),
	}))

	if err != nil {
		return nil, err
	}

	return &Service{db: db}, nil
}

func GetDB() *sqldb.Database {
	return myDB
}
