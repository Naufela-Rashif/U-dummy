package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var DB *sql.DB

//func Connect() (*sql.DB, error) {
//	// Memuat konfigurasi database
//	dbConfig, err := config.LoadDatabaseConfig()
//	if err != nil {
//		log.Fatalf("Error loading database config: %v", err)
//	}
//
//	// Membuat string koneksi (DSN) untuk PostgreSQL
//	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
//		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password,
//		dbConfig.Name, dbConfig.SSLMode)
//
//	// Membuka koneksi ke database
//	db, err := sql.Open("postgres", dsn)
//	if err != nil {
//		log.Fatalf("Error while opening the database: %v", err)
//	}
//	defer func(db *sql.DB) {
//		err := db.Close()
//		if err != nil {
//			log.Fatalf("Error while closing the database: %v", err)
//		}
//	}(db)
//
//	// Melakukan ping untuk memastikan koneksi berhasil
//	err = db.Ping()
//	if err != nil {
//		log.Fatalf("Failed to ping the database: %v", err)
//	}
//
//	// Jika berhasil, tampilkan pesan
//	fmt.Println("Successfully connected to the database!")
//	return db, nil
//}

func Connect() {
	dsn := "host=localhost user=udummy password=kujwhehfuw32433 dbname=db_udummy port=5432 sslmode=disable"
	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Error opening database connection:", err)
	}

	// Test database connection
	err = DB.Ping()
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	log.Println("Successfully connected to the database!")
}
