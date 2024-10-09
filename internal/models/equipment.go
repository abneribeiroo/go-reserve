package models

import (
	"database/sql"
	"time"
)

type Equipment struct {
    ID          int       `json:"id"`
    Name        string    `json:"name"`
    Description string    `json:"description"`
    Quantity    int       `json:"quantity"`
    CreatedAt   time.Time `json:"created_at"`
    OwnerID     int    `json:"owner_id"`

}

func (e *Equipment) Create(db *sql.DB) error {
    _, err := db.Exec("INSERT INTO equipment (name, description, quantity, created_at, owner_id) VALUES ($1, $2, $3, $4, $5)", e.Name, e.Description, e.Quantity, time.Now(), e.OwnerID)
    return err
}