package model

type Townshipe struct {
	ID         int        `json:"id"`
	Name       string     `json:"name"`
	Department Department `json:"department"`
}
