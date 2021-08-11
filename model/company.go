package model

type Comapny struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Township Township `json:"township"`
}
