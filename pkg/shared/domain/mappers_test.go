package domain_test

import (
	authDomain "ephelsa/my-career/pkg/auth/domain"
	"ephelsa/my-career/pkg/shared/domain"
	"github.com/magiconair/properties/assert"
	"testing"
)

func Test_AuthDomainRegisterToUserDomainUser(t *testing.T) {
	register := authDomain.Register{
		Email:             "xephelsax@gmail.com",
		DocumentType:      "CC",
		Document:          "123123",
		FirstName:         "Leonardo",
		SecondName:        "Andres",
		FirstSurname:      "Perez",
		SecondSurname:     "Castilla",
		Password:          "123123123",
		InstitutionType:   1,
		InstitutionName:   "University of Antioquia",
		StudyLevel:        4,
		RegistryConfirmed: true,
		CountryCode:       "CO",
		DepartmentCode:    "70",
		MunicipalityCode:  "001",
		Birthdate:         "10-05-1997",
	}

	got := domain.AuthDomainRegisterToUserDomainUser(register)

	assert.Equal(t, got.FirstName, register.FirstName)
	assert.Equal(t, got.SecondName, register.SecondName)
	assert.Equal(t, got.FirstSurname, register.FirstSurname)
	assert.Equal(t, got.SecondSurname, register.SecondSurname)
	assert.Equal(t, got.Email, register.Email)
	assert.Equal(t, got.CountryCode, register.CountryCode)
	assert.Equal(t, got.DepartmentCode, register.DepartmentCode)
	assert.Equal(t, got.MunicipalityCode, register.MunicipalityCode)
	assert.Equal(t, got.DocumentTypeCode, register.DocumentType)
	assert.Equal(t, got.Document, register.Document)
	assert.Equal(t, got.InstitutionTypeCode, register.InstitutionType)
	assert.Equal(t, got.InstitutionName, register.InstitutionName)
	assert.Equal(t, got.StudyLevelCode, register.StudyLevel)
	assert.Equal(t, got.Birthdate, register.Birthdate)
}
