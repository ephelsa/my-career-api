package domain

type Country struct {
	ISOCode string `json:"iso_code"`
	Name    string `json:"name"`
}

type Department struct {
	CountryCode    string `json:"country_code"`
	DepartmentCode string `json:"department_code"`
	Name           string `json:"name"`
}

type Municipality struct {
	CountryCode      string `json:"country_code"`
	DepartmentCode   string `json:"department_code"`
	MunicipalityCode string `json:"municipality_code"`
	Name             string `json:"name"`
}
