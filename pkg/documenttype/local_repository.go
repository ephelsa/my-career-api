package documenttype

import (
	"context"
	"ephelsa/my-career/internal/database"
)

type LocalRepository struct {
	Data *database.Data
}

func (lr *LocalRepository) FetchAll(ctx context.Context) (result []DocumentType, err error) {
	query := `SELECT id, value FROM document_type;`

	rows, err := lr.Data.Database.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer func() {
		err = rows.Close()
	}()

	for rows.Next() {
		var r DocumentType
		err = rows.Scan(&r.ID, &r.Name)
		if err != nil {
			return nil, err
		}

		result = append(result, r)
	}

	return result, nil
}

func (lr *LocalRepository) New(ctx context.Context, dt *DocumentType) (err error) {
	query := `INSERT INTO document_type (id, value) VALUES($1, $2) RETURNING id;`

	stmt, err := lr.Data.Database.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	defer func() {
		err = stmt.Close()
	}()

	row := stmt.QueryRowContext(ctx, dt.ID, dt.Name)

	err = row.Scan(&dt.ID)
	if err != nil {
		return err
	}

	return nil
}
