package models

import (
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
