package models

import "time"

// ExperienceOutputModel is the structure of the data returned while retrieving
// details of experience.
type ExperienceOutputModel struct {
	ID          int32     `db:"id"`
	Title       string    `db:"title"`
	CompanyName string    `db:"company_name"`
	StartDate   time.Time `db:"start_date"`
	EndDate     time.Time `db:"end_date"`
	IsCurrent   bool      `db:"is_current"`
	Location    string    `db:"location"`
	Description []string  `db:"description"`
	Skills      []string  `db:"skills"`
}
