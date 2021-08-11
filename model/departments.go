package model

type Departments struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Region Region `json:"region"`
}
