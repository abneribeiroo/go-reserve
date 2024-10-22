package models

import (
	"database/sql"
	"errors"
	"time"
)

type Reservation struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	EquipmentID int       `json:"equipment_id"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}

const (
	StatusPending  = "pending"
	StatusApproved = "approved"
	StatusRejected = "rejected"
)

func NewReservation(userID, equipmentID int, startTime, endTime time.Time) *Reservation {
	return &Reservation{
		UserID:      userID,
		EquipmentID: equipmentID,
		StartTime:   startTime,
		EndTime:     endTime,
		Status:      StatusPending,
		CreatedAt:   time.Now(),
	}
}

// Função para salvar a reserva no banco de dados
func (r *Reservation) Create(db *sql.DB) error {
	if r.StartTime.After(r.EndTime) {
		return errors.New("start time cannot be after end time")
	}

	// Verifica se há conflito com outras reservas aprovadas
	if hasConflict, err := r.hasTimeConflict(db); err != nil {
		return err
	} else if hasConflict {
		return errors.New("time conflict with another approved reservation")
	}

	// Insere a reserva no banco
	_, err := db.Exec("INSERT INTO reservations (user_id, equipment_id, start_time, end_time, status, created_at) VALUES ($1, $2, $3, $4, $5, $6)", r.UserID, r.EquipmentID, r.StartTime, r.EndTime, r.Status, r.CreatedAt)
	return err
}

// Método para verificar conflitos de horário
func (r *Reservation) hasTimeConflict(db *sql.DB) (bool, error) {
	var count int
	query := `
		SELECT COUNT(*) 
		FROM reservations 
		WHERE equipment_id = $1 
		AND status = $2 
		AND ($3 < end_time AND $4 > start_time)
	`
	err := db.QueryRow(query, r.EquipmentID, StatusApproved, r.StartTime, r.EndTime).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// Método para aprovar uma reserva
func (r *Reservation) Approve(db *sql.DB) error {
	if r.Status != StatusPending {
		return errors.New("reservation is not in pending status")
	}

	// Verifica conflitos de horário
	if hasConflict, err := r.hasTimeConflict(db); err != nil {
		return err
	} else if hasConflict {
		return errors.New("cannot approve reservation due to time conflict")
	}

	// Atualiza o status da reserva para aprovada
	r.Status = StatusApproved
	_, err := db.Exec("UPDATE reservations SET status = $1 WHERE id = $2", r.Status, r.ID)
	return err
}

// Método para rejeitar uma reserva
func (r *Reservation) Reject(db *sql.DB) error {
	if r.Status != StatusPending {
		return errors.New("reservation is not in pending status")
	}

	// Atualiza o status da reserva para rejeitada
	r.Status = StatusRejected
	_, err := db.Exec("UPDATE reservations SET status = $1 WHERE id = $2", r.Status, r.ID)
	return err
}
