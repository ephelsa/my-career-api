package data

import (
	"database/sql"
	sharedServer "ephelsa/my-career/pkg/shared/infrastructure/server"

	documentTypeDatabase "ephelsa/my-career/pkg/documenttype/infraestructure/database"
	documentTypeServer "ephelsa/my-career/pkg/documenttype/infraestructure/server"

	studyLevelDatabase "ephelsa/my-career/pkg/studylevel/infraestructure/database"
	studyLevelServer "ephelsa/my-career/pkg/studylevel/infraestructure/server"

	institutionTypeDatabase "ephelsa/my-career/pkg/institutiontype/infraestructure/database"
	institutionTypeServer "ephelsa/my-career/pkg/institutiontype/infraestructure/server"
)

func AddServerRouter(s *sharedServer.Server, db *sql.DB) {
	documentTypeServer.NewDocumentTypeServer(s.Server, documentTypeDatabase.NewPostgresDocumentTypeRepository(db))
	studyLevelServer.NewStudyLevelServer(s.Server, studyLevelDatabase.NewPostgresStudyLevelRepository(db))
	institutionTypeServer.NewInstitutionTypeServer(s.Server, institutionTypeDatabase.NewPostgresInstitutionTypeRepository(db))
}
