package models

import (
	"time"
)

type Movie struct {
	ID          int            `json:"id"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Year        int            `json:"year"`
	ReleaseDate time.Time      `json:"releaseDate"`
	Runtime     int            `json:"runtime"`
	Rating      int            `json:"rating"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	MovieGenre  []MovieGenre   `json:"-"`
}

type Genre struct {
	ID          int            `json:"id"`
	GenreName   string         `json:"genreName"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
}

type MovieGenre struct {
	ID          int            `json:"id"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	MovieID     int            `json:"movieID"`
	GenreID     int            `json:"genreID"`
}
