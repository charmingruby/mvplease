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
			(name, email, role, password, groups_quantity)
			VALUES ($1, $2, $3, $4, $5)
			RETURNING *`,
		saveAccount: `UPDATE accounts
			SET name = $1, email = $2, role = $3, avatar_url = $4, password = $5, groups_quantity = $6, deleted_by = $7, updated_at = $8, deleted_at = $9 
			WHERE id = $10 AND deleted_at IS NULL
			RETURNING *`,
		deleteAccount: `UPDATE accounts
			SET deleted_by = $1, deleted_at = $2, updated_at = $3
			WHERE id = $4 AND deleted_at IS NULL
			RETURNING *`,
	}
}
