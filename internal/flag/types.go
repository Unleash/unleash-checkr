package flag

import "time"

// Flags struct
type Flags struct {
	Features []Flag `json:"features"`
}

// Flag struct
type Flag struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	IsEnabled   bool      `json:"enabled"`
}
