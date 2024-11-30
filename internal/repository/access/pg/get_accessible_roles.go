package access

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/marinaaaniram/go-common-platform/pkg/db"

	"github.com/marinaaaniram/go-auth/internal/errors"
)

// GetAccessibleRoles Access in repository layer
func (r *repo) GetAccessibleRoles(ctx context.Context, endpointAddress string) ([]string, error) {
	builder := sq.Select(roleColumn).
		From(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{endpointColumn: endpointAddress})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, errors.ErrFailedToBuildQuery(err)
	}

	q := db.Query{
		Name:     "access_repository.GetAccessibleRoles",
		QueryRaw: query,
	}

	rows, err := r.db.DB().QueryContext(ctx, q, args...)
	if err != nil {
		return nil, errors.ErrFailedToSelectQuery(err)
	}
	defer rows.Close()

	var roles []string
	for rows.Next() {
		var role string
		if err := rows.Scan(&role); err != nil {
			return nil, errors.ErrFailedToScanRow(err)
		}
		roles = append(roles, role)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return roles, nil
}
