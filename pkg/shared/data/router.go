package data

import (
	"database/sql"
	"ephelsa/my-career/internal/env"
	sharedServer "ephelsa/my-career/pkg/shared/infrastructure/server"

	surveyDatabase "ephelsa/my-career/pkg/survey/infraestructure/database"
	surveyServer "ephelsa/my-career/pkg/survey/infraestructure/server"

	documentTypeDatabase "ephelsa/my-career/pkg/documenttype/infraestructure/database"
	documentTypeServer "ephelsa/my-career/pkg/documenttype/infraestructure/server"

	studyLevelDatabase "ephelsa/my-career/pkg/studylevel/infraestructure/database"
	studyLevelServer "ephelsa/my-career/pkg/studylevel/infraestructure/server"

	institutionTypeDatabase "ephelsa/my-career/pkg/institutiontype/infraestructure/database"
	institutionTypeServer "ephelsa/my-career/pkg/institutiontype/infraestructure/server"

	locationDatabase "ephelsa/my-career/pkg/location/infraestructure/database"
	locationServer "ephelsa/my-career/pkg/location/infraestructure/server"

	authDatabase "ephelsa/my-career/pkg/auth/infraestructure/database"
	authServer "ephelsa/my-career/pkg/auth/infraestructure/server"

	userDatabase "ephelsa/my-career/pkg/user/infraestructure/database"
	userServer "ephelsa/my-career/pkg/user/infraestructure/server"
)

func ServerRouter(s *sharedServer.Server, db *sql.DB, model env.ClassifierModel) {
	userPostgresRepo := userDatabase.NewPostgresUserRepository(db)

	documentTypeServer.NewDocumentTypeServer(s.Server, documentTypeDatabase.NewPostgresDocumentTypeRepository(db))
	studyLevelServer.NewStudyLevelServer(s.Server, studyLevelDatabase.NewPostgresStudyLevelRepository(db))
	institutionTypeServer.NewInstitutionTypeServer(s.Server, institutionTypeDatabase.NewPostgresInstitutionTypeRepository(db))
	locationServer.NewLocationServer(s.Server, locationDatabase.NewPostgresLocationRepository(db))
	authServer.NewAuthServer(s.Server, authDatabase.NewPostgresAuthDatabase(db), userPostgresRepo)
	surveyServer.NewSurveyServer(s.Server, surveyDatabase.NewPostgresSurveyRepository(db), model)
	userServer.NewUserServer(s.Server, userPostgresRepo)
}
