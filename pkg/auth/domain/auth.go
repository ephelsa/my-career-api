package domain

const PasswordMinLen = 8

type Register struct {
	Email        string `json:"email"`
	DocumentType string `json:"document_type"`
	Document     string `json:"document"`

	FirstName     string `json:"first_name"`
	SecondName    string `json:"second_name"`
	FirstSurname  string `json:"first_surname"`
	SecondSurname string `json:"second_surname"`

	Password string `json:"password"`

	InstitutionType int    `json:"institution_type"`
	InstitutionName string `json:"institution_name"`

	StudyLevel int `json:"study_level"`

	RegistryConfirmed bool `json:"registry_confirmed"`

	CountryCode      string `json:"country_code"`
	DepartmentCode   string `json:"department_code"`
	MunicipalityCode string `json:"municipality_code"`
}

type RegisterSuccess struct {
	Email string `json:"email"`
}

type LoginSuccess struct {
	Token string `json:"token"`
}
