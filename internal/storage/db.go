package storage

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/arldka/github-releases-gchat-notifier/internal/models"
)

type DB struct {
	*gorm.DB
}

func NewDB() DB {
	fmt.Println("Start of NewDB")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = "26257"
	}
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=verify-full", dbUser, dbPassword, dbHost, dbPort, dbName)
	fmt.Println("After dsn: " + dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if db == nil {
		fmt.Println("db is null")
	}
	fmt.Println("After gorm.Open")
	if err != nil {
		log.Fatal("failed to connect database", err)
	}

	fmt.Println("After error Handler connect to database")

	// Create the repositories table if it does not exist.
	fmt.Println("Before Create if not exists")
	/* 	if err := db.Exec(`CREATE TABLE IF NOT EXISTS repositories (
		id SERIAL PRIMARY KEY,
		name STRING NOT NULL,
		tag STRING NOT NULL,
		notified BOOLEAN NOT NULL
		)`); err != nil {
		fmt.Println("Error in create table if not exists: ")
		fmt.Println(err)
		log.Fatal(err)
	} */

	fmt.Println("Before Return")

	return DB{db}
}

func (db DB) ListRepositories() *[]models.Repository {
	var repositories []models.Repository
	db.Find(&repositories)

	return &repositories
}

func (db DB) UpdateRepositories(repositories *[]models.Repository) {
	for _, r := range *repositories {
		db.Save(&r)
	}
}
