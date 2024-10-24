package config

import (
	"database/sql"
	"fmt"
	"golang-porto/backend/pkg/models"
	"log"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() *gorm.DB {
	var err error

	// Nama database dan DSN
	dbName := "porto"
	dsnWithoutDB := "host=localhost user=postgres password=cimapag1 port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	dsnWithDB := fmt.Sprintf("%s dbname=%s", dsnWithoutDB, dbName)

	// Cek dan buat database jika tidak ada
	err = createDatabaseIfNotExists(dbName, dsnWithoutDB)
	if err != nil {
		log.Fatalf("Failed to create database: %v", err)
	}

	// Membuka koneksi ke database PostgreSQL
	DB, err = gorm.Open(postgres.Open(dsnWithDB), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Mendapatkan objek database SQL
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}

	// Melakukan ping untuk mengecek koneksi
	err = sqlDB.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	log.Println("Successfully connected to the PostgreSQL database.")

	// Panggil fungsi untuk melakukan migrasi
	runMigrations()

	return DB
}

func createDatabaseIfNotExists(dbName string, dsn string) error {
	// Membuka koneksi ke PostgreSQL tanpa database
	tempDB, err := sql.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("failed to open database connection: %v", err)
	}
	defer tempDB.Close()

	// Mengecek apakah database sudah ada
	var exists bool
	err = tempDB.QueryRow("SELECT EXISTS(SELECT datname FROM pg_catalog.pg_database WHERE datname = $1)", dbName).Scan(&exists)
	if err != nil {
		return fmt.Errorf("failed to check if database exists: %v", err)
	}

	// Jika tidak ada, buat database
	if !exists {
		_, err = tempDB.Exec(fmt.Sprintf("CREATE DATABASE %s", dbName))
		if err != nil {
			return fmt.Errorf("failed to create database: %v", err)
		}
		log.Printf("Database '%s' created successfully.\n", dbName)
	}

	return nil
}

// Fungsi untuk menjalankan migrasi ulang jika diperlukan
func runMigrations() {
	log.Println("Checking and running database migrations...")

	// AutoMigrate akan memeriksa model dan melakukan perubahan jika diperlukan
	err := DB.AutoMigrate(
		&models.User{},
		&models.Profile{},
		&models.Project{},
		&models.Skill{},
		&models.Contact{},
		&models.Education{},
		&models.Experience{},
	)

	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Println("Database migration completed successfully.")
}

func ConnectDatabaseSqlite() {
	var err error
	DB, err = gorm.Open(sqlite.Open("porto.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Menjalankan query sederhana untuk mengecek koneksi
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}

	err = sqlDB.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	log.Println("Successfully connected to the database.")

	// Migrasi model
	err = DB.AutoMigrate(&models.User{}, &models.Project{}, &models.Skill{}, &models.Contact{}, &models.Education{}, &models.Experience{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}
