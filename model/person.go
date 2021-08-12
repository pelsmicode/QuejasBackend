package model

type Person struct {
	ID       int         `json:"id"`
	Name     string      `json:"name"`
	Lastname string      `json:"lastname"`
	DPI      string      `json:"dpi"`
	NIT      string      `json:"nit"`
	Gender   string      `json:"gender"`
	Email    string      `json:"email"`
	Phone    string      `json:"phone"`
	Towship  Township    `json:"towship"`
	Branch   DiacoBranch `json:"branch"`
}
