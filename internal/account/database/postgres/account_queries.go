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
			(name, email, role, password)
			VALUES ($1, $2, $3, $4)
			RETURNING *`,
		saveAccount: `UPDATE accounts
			SET name = $1, email = $2, role = $3, avatar_url = $4, password = $5, deleted_by = $6, updated_at = $7, deleted_at = $8 
			WHERE id = $9 AND deleted_at IS NULL
			RETURNING *`,
		deleteAccount: `UPDATE accounts
			SET deleted_by = $1, deleted_at = $2, updated_at = $3
			WHERE id = $4 AND deleted_at IS NULL
			RETURNING *`,
	}
}
