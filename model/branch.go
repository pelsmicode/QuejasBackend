package model

type DiacoBranch struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Township Township `json:"township"`
}
