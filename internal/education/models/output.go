package models

import "time"

// EducationOutputModel is the data structure that represents the format of the
// education data when returned by a query.
type EducationOutputModel struct {
	ID        int32     `db:"id"`
	Institute string    `db:"institute"`
	Degree    string    `db:"degree"`
	Major     string    `db:"major"`
	Grade     string    `db:"grade"`
	Location  string    `db:"location"`
	StartDate time.Time `db:"start_date"`
	EndDate   time.Time `db:"end_date"`
}
