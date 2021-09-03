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
	Township Township    `json:"township" db:"township"`
	Branch   DiacoBranch `json:"branch" db:"branch"`
}

type PersonResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	DPI      string `json:"dpi"`
	NIT      string `json:"nit"`
	Gender   string `json:"gender"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type PersonRequest struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	DPI      string `json:"dpi"`
	NIT      string `json:"nit"`
	Gender   string `json:"gender"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Township int    `json:"township"`
	Branch   int    `json:"branch"`
}

func ToPresonResponse(p PersonRequest) PersonResponse {
	return PersonResponse{
		ID:       p.ID,
		Name:     p.Name,
		Lastname: p.Lastname,
		DPI:      p.DPI,
		NIT:      p.NIT,
		Gender:   p.Gender,
		Email:    p.Email,
		Phone:    p.Phone,
	}
}
