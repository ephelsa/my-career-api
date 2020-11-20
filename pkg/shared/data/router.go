package data

import (
	"database/sql"
	documentTypeDatabase "ephelsa/my-career/pkg/documenttype/infraestructure/database"
	documentTypeServer "ephelsa/my-career/pkg/documenttype/infraestructure/server"
	sharedServer "ephelsa/my-career/pkg/shared/infrastructure/server"
)

func AddServerRouter(s *sharedServer.Server, db *sql.DB) {
	documentTypeServer.NewDocumentTypeServer(s.Server, documentTypeDatabase.NewPostgresDocumentTypeRepository(db))
}
