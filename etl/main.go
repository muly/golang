package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func initDB() *sql.DB {
	//TODO: return error

	connStr := "user=postgres password=postgres dbname=postgres sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping the database: %v", err)
	}

	fmt.Println("Successfully connected to the PostgreSQL database!")
	return db
}

type etl[Stage any, Regular any] interface {
	Extract(tx *sql.Tx) ([]Stage, error)       // Extract data from the stage table
	Transform(data []Stage) ([]Regular, error) // Transform stage data into main data
	Load(data []Regular, tx *sql.Tx) error     // Load regular data into the main table
}

func main() {
	db := initDB()
	defer db.Close()

	tx, err := db.Begin() // Start a shared transaction
	if err != nil {
		log.Fatalf("Failed to begin transaction: %v", err)
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback() // Rollback on panic
			panic(p)
		} else if err != nil {
			tx.Rollback() // Rollback on error
		} else {
			err = tx.Commit() // Commit if no error
		}
	}()

	// Run ETL for students
	if err := runETL(&studentETL{db: db}, tx); err != nil {
		log.Fatalf("Student ETL process failed: %v", err)
	}

	// Run ETL for schools
	if err := runETL(&schoolETL{db: db}, tx); err != nil {
		log.Fatalf("School ETL process failed: %v", err)
	}
}

func runETL[Stage any, Regular any](etlProcess etl[Stage, Regular], tx *sql.Tx) error {
	stage, err := etlProcess.Extract(tx)
	if err != nil {
		return fmt.Errorf("failed to extract data: %v", err)
	}

	// Transform the data
	maindata, err := etlProcess.Transform(stage)
	if err != nil {
		return fmt.Errorf("failed to transform data: %v", err)
	}

	// Load the data into the target table
	err = etlProcess.Load(maindata, tx)
	if err != nil {
		return fmt.Errorf("failed to load data: %v", err)
	}

	log.Println("ETL process completed successfully!")
	return nil
}
