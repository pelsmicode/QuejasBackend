package model

type Township struct {
	ID         int        `json:"id"`
	Name       string     `json:"name"`
	Department Department `json:"department" db:"department"`
}
