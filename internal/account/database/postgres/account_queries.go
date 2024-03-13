package postgres

const (
	getAccountByID    = "get account by id"
	getAccountByEmail = "get account by email"
	fetchAccounts     = "fetch accounts"
	createAccount     = "create account"
	saveAccount       = "save account"
	deleteAccount     = "delete account"
)

func queriesAccount() map[string]string {
	return map[string]string{
		getAccountByID: `SELECT * FROM accounts 
			WHERE id = $1`,
		getAccountByEmail: `SELECT * FROM accounts 
			WHERE email = $1`,
		fetchAccounts: `SELECT * FROM accounts 
			LIMIT $1 
			OFFSET $2`,
		createAccount: `INSERT INTO accounts 
			(name, email, role, password, aggregates_quantity, examples_quantity)
			VALUES ($1, $2, $3, $4, $5, $6)
			RETURNING *`,
		saveAccount: `UPDATE accounts
			SET name = $1, email = $2, role = $3, avatar_url = $4, password = $5, aggregates_quantity = $6, examples_quantity = $7, deleted_by = $8, updated_at = $9, deleted_at = $10 
			WHERE id = $11 AND deleted_at IS NULL
			RETURNING *`,
		deleteAccount: `UPDATE accounts
			SET deleted_by = $1, deleted_at = $2 
			WHERE id = $3 AND deleted_at IS NULL
			RETURNING *`,
	}
}
