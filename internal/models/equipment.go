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
	_, err := db.Exec("INSERT INTO equipment (name, description, quantity, created_at, owner_id) VALUES ($1, $2, $3, $4, $5)", 
		e.Name, e.Description, e.Quantity, time.Now(), e.OwnerID)
	return err
}

func GetAllEquipment(db *sql.DB) ([]Equipment, error) {
	rows, err := db.Query("SELECT id, name, description, quantity, created_at, owner_id FROM equipment")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var equipmentList []Equipment
	for rows.Next() {
		var equipment Equipment
		if err := rows.Scan(&equipment.ID, &equipment.Name, &equipment.Description, &equipment.Quantity, &equipment.CreatedAt, &equipment.OwnerID); err != nil {
			return nil, err
		}
		equipmentList = append(equipmentList, equipment)
	}
	return equipmentList, nil
}


func GetEquipmentById(db *sql.DB, id int) (*Equipment, error) {
	var equipment Equipment
	err := db.QueryRow("SELECT id, name, description, quantity, created_at, owner_id FROM equipment WHERE id = $1", id).Scan(
		&equipment.ID, &equipment.Name, &equipment.Description, &equipment.Quantity, &equipment.CreatedAt, &equipment.OwnerID)
	if err != nil {
		return nil, err
	}
	return &equipment, nil
}


func (e *Equipment) Update(db *sql.DB) error {
	_, err := db.Exec("UPDATE equipment SET name = $1, description = $2, quantity = $3 WHERE id = $4",
		e.Name, e.Description, e.Quantity, e.ID)
	return err
}

func DeleteEquipment(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM equipment WHERE id = $1", id)
	return err
}