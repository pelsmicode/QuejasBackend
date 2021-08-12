package model

import "time"

type Complaint struct {
	ID        int       `json:"id"`
	NoDoc     string    `json:"nodoc"`
	DateDoc   string    `json:"datedoc"`
	Detail    string    `json:"detail"`
	Petition  string    `json:"petition"`
	CreatedAt time.Time `json:"created_at"`
	Company   Comapny   `json:"company"`
	Person    Person    `json:"person"`
}
