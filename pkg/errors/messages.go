package errors

import "fmt"

func StatementPreparationError(queryName string) string {
	return fmt.Sprintf("error preparing statement: %s", queryName)
}

func StatementNotPreparedError(queryName string) string {
	return fmt.Sprintf("prepared statement %s not found", queryName)

}

func NotFoundError(entity string) string {
	return fmt.Sprintf("%s not found", entity)
}
