package models

// ProjectOutputModel represents the structure of data returned while getting
// details of a project.
type ProjectOutputModel struct {
	ID           int32    `db:"id"`
	Name         string   `db:"name"`
	PublicURL    string   `db:"public_url"`
	RepoURL      string   `db:"repo_url"`
	Description  []string `db:"description"`
	Technologies []string `db:"technologies"`
}
