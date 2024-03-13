package errors

import "fmt"

func StatementPreparationError(queryName string) string {
	return fmt.Sprintf("error preparing statement: %s", queryName)
}

func StatementNotPrepared(queryName string) string {
	return fmt.Sprintf("prepared statement %s not found", queryName)

}
