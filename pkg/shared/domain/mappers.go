package domain

import (
	authDomain "ephelsa/my-career/pkg/auth/domain"
	userDomain "ephelsa/my-career/pkg/user/domain"
)

// AuthDomainRegisterToUserDomainUser map authDomain.Register to userDomain.User
func AuthDomainRegisterToUserDomainUser(r authDomain.Register) userDomain.User {
	return userDomain.User{
		FirstName:           r.FirstName,
		SecondName:          r.SecondName,
		FirstSurname:        r.FirstSurname,
		SecondSurname:       r.SecondSurname,
		Email:               r.Email,
		CountryCode:         r.CountryCode,
		DepartmentCode:      r.DepartmentCode,
		MunicipalityCode:    r.MunicipalityCode,
		DocumentTypeCode:    r.DocumentType,
		Document:            r.Document,
		InstitutionTypeCode: r.InstitutionType,
		InstitutionName:     r.InstitutionName,
		StudyLevelCode:      r.StudyLevel,
		Birthdate:           r.Birthdate,
	}
}
