package models

// ProjectInputModel is the data structure used to create a new project record.
type ProjectInputModel struct {
	Name         string   `db:"name"`
	PublicURL    string   `db:"public_url"`
	RepoURL      string   `db:"repo_url"`
	Description  []string `db:"description"`
	Technologies []string `db:"technologies"`
}
