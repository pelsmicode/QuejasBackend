package model

type Comapny struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	NIT      string   `json:"nit"`
	Phone    string   `json:"phone"`
	Email    string   `json:"email"`
	Addres   string   `json:"address"`
	Township Township `json:"township" db:"township"`
}

type CompanyRequest struct {
	ID       int    `json:"id"`
	Name     string `json:"name" db:"name"`
	NIT      string `json:"nit" db:"nit"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Addres   string `json:"address" db:"address"`
	Township string `json:"township" db:"township_id"`
}

type ComapnyResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	NIT      string `json:"nit"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Addres   string `json:"address"`
	Township int    `json:"township" db:"township_id"`
}

func ToCompanyReponse(c CompanyRequest) ComapnyResponse {
	return ComapnyResponse{
		ID:     c.ID,
		Name:   c.Name,
		NIT:    c.NIT,
		Phone:  c.NIT,
		Email:  c.Email,
		Addres: c.Addres,
	}
}
