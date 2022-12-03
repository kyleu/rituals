package comment

import (
	"context"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/kyleu/rituals/app/lib/database"
	"github.com/kyleu/rituals/app/util"
)

func (s *Service) GetByModels(ctx context.Context, tx *sqlx.Tx, logger util.Logger, vals ...any) (Comments, error) {
	if len(vals) == 0 {
		return nil, nil
	}
	if len(vals)%2 != 0 {
		return nil, errors.New("must provide even number of arguments")
	}
	size := len(vals) / 2
	wc := make([]string, 0, size)
	for i := 0; i < len(vals); i += 2 {
		wc = append(wc, fmt.Sprintf(`("svc" = $%d and "model_id" = $%d)`, i+1, i+2))
	}
	q := database.SQLSelectSimple(columnsString, tableQuoted, strings.Join(wc, " or "))
	return s.ListSQL(ctx, tx, q, logger, vals...)
}
