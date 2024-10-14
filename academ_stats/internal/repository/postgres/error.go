package postgres

import (
	"fmt"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

func customErr(message string, err error) error {
	// postgres errors
	if pgErr, ok := err.(*pgconn.PgError); ok {
		// duplicate key
		// if pgErr.Code == pgerrcode.UniqueViolation {
		// 	return response.ErrDuplicateKey
		// }
		// // unknown keys
		// if pgErr.Code == pgerrcode.ForeignKeyViolation {
		// 	return response.ErrForeignKey
		// }
		// custom errors from trigger
		if pgErr.Code == pgerrcode.RaiseException {
			// if pgErr.Message == response.ErrEndStartDate.Error() {
			// 	return response.ErrEndStartDate
			// }
			// if pgErr.Message == response.ErrEndEndDate.Error() {
			// 	return response.ErrEndEndDate
			// }
		}
	}
	// service errors
	if message != "" {
		return fmt.Errorf("%s: %s", message, err)
	}
	return err
}
