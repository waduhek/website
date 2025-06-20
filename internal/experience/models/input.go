package models

import "time"

// ExperienceInputModel is the structure of the data required to create a new
// entry for experience.
type ExperienceInputModel struct {
	Title       string    `db:"title"`
	CompanyName string    `db:"company_name"`
	StartDate   time.Time `db:"start_date"`
	EndDate     time.Time `db:"end_date"`
	IsCurrent   bool      `db:"is_current"`
	Location    string    `db:"location"`
	Description []string  `db:"description"`
	Skills      []string  `db:"skills"`
}
