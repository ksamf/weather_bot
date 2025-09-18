package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

type History struct {
	ID          int
	UserId      int64
	City        string
	Temperature float32
	Description string
	CreatedAt   time.Time
}

var DB *sql.DB

func InitDb() {
	pgInfo := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("PG_USER"), os.Getenv("PG_PASS"), os.Getenv("PG_HOST"),
		os.Getenv("PG_PORT"), os.Getenv("PG_NAME"))
	var err error
	DB, err = sql.Open("postgres", pgInfo)
	if err != nil {
		panic(err)
	}
	_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS history (
    id SERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    city TEXT NOT NULL,
    temperature REAL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`)
	if err != nil {
		log.Fatalf("Error create database:%v", err)
	}
}

func GetHistory(id int64, limit int) []History {
	rows, err := DB.Query("SELECT city, temperature, description, created_at FROM history WHERE user_id=$1 ORDER BY created_at DESC LIMIT $2", id, limit)

	if err != nil {
		log.Printf("Error get history: %v", err)
	}
	defer rows.Close()
	var history []History
	for rows.Next() {
		var h History
		if err := rows.Scan(&h.City, &h.Temperature, &h.Description, &h.CreatedAt); err != nil {
			log.Printf("Error scan rows: %v", err)
		}
		history = append(history, h)
	}

	return history
}
func InsertHistory(id int64, name string, temp float32, description string) {
	_, err := DB.Exec(
		"INSERT INTO history (user_id, city, temperature, description) VALUES($1,$2,$3,$4)",
		id, name, temp, description)

	if err != nil {
		log.Printf("Error insert in table history:%v", err)
	}
}
