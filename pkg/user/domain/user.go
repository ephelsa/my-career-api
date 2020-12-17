package domain

type User struct {
	FirstName     string `json:"first_name"`
	SecondName    string `json:"second_name"`
	FirstSurname  string `json:"first_surname"`
	SecondSurname string `json:"second_surname"`

	Email string `json:"email"`

	CountryCode      string `json:"country_code,omitempty"`
	Country          string `json:"country,omitempty"`
	DepartmentCode   string `json:"department_code,omitempty"`
	Department       string `json:"department,omitempty"`
	MunicipalityCode string `json:"municipality_code,omitempty"`
	Municipality     string `json:"municipality,omitempty"`

	DocumentTypeCode string `json:"document_type_code,omitempty"`
	DocumentType     string `json:"document_type,omitempty"`
	Document         string `json:"document,omitempty"`

	InstitutionTypeCode int    `json:"institution_type_code,omitempty"`
	InstitutionType     string `json:"institution_type,omitempty"`
	InstitutionName     string `json:"institution_name,omitempty"`

	StudyLevelCode int    `json:"study_level_code,omitempty"`
	StudyLevel     string `json:"study_level,omitempty"`

	Birthdate string `json:"birthdate,omitempty"`
}
