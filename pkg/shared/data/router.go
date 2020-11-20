package data

import (
	"database/sql"
	documentTypeDatabase "ephelsa/my-career/pkg/documenttype/infraestructure/database"
	documentTypeServer "ephelsa/my-career/pkg/documenttype/infraestructure/server"
	sharedServer "ephelsa/my-career/pkg/shared/infrastructure/server"

	studyLevelDatabase "ephelsa/my-career/pkg/studylevel/infraestructure/database"
	studyLevelServer "ephelsa/my-career/pkg/studylevel/infraestructure/server"
)

func AddServerRouter(s *sharedServer.Server, db *sql.DB) {
	documentTypeServer.NewDocumentTypeServer(s.Server, documentTypeDatabase.NewPostgresDocumentTypeRepository(db))
	studyLevelServer.NewStudyLevelServer(s.Server, studyLevelDatabase.NewPostgresStudyLevelRepository(db))
}
