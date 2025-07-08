package models

import "time"

// EducationInputModel is the data structure used for creating a new education
// record.
type EducationInputModel struct {
	Institute string    `db:"institute"`
	Degree    string    `db:"degree"`
	Major     string    `db:"major"`
	Grade     string    `db:"grade"`
	Location  string    `db:"location"`
	StartDate time.Time `db:"start_date"`
	EndDate   time.Time `db:"end_date"`
}
