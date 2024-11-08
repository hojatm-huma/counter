package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type Database struct {
	Connection *gorm.DB
}

func NewDatabase() (*Database, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s ",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	db, err := gorm.Open(postgres.Open(dsn), nil)
	if err != nil {
		return &Database{}, err
	}

	return &Database{db}, nil
}

func (me *Database) Migrate() {
	me.Connection.AutoMigrate(&Counter{})
}

type Counter struct {
	gorm.Model
	Value uint
}

func (me *Counter) Increment() {
	me.Value++
}

func (db *Database) FirstOrCreateCounter() Counter {
	counter := Counter{}
	db.Connection.FirstOrCreate(&counter, Counter{Value: 0})

	return counter
}

func (db *Database) SaveCounter(counter Counter) {
	db.Connection.Save(&counter)
}
