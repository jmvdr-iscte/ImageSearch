package models

import (
	"time"

	"github.com/google/uuid"
)

type Image struct {
	UID       uuid.UUID `json:"uid"`
	Name      string    `json:"name"`
	Path      string    `json:"path"`
	Width     int       `json:"width"`
	Height    int       `json:"height"`
	CreatedAt time.Time `json:"created_at"`
}
