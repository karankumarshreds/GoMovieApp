package models 

import "time"

type MovieGenre struct {
	ID          int            `json:"id"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	MovieID     int            `json:"movieID"`
	GenreID     int            `json:"genreID"`
}
