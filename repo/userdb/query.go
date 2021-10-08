package userdb

const (
	QueryUser = `
		SELECT
			id,
			name,
			email,
			created_at,
			updated_at
		FROM
			users
	`

	QueryGetUser = `
		SELECT
			id,
			name,
			email,
			created_at,
			updated_at
		FROM
			users
		WHERE
			id = ?;
	`

	QueryCreateUser = `
		INSERT INTO
			users
		(
			name,
			email,
			created_at,
			updated_at
		)
		VALUES 
		(
			'%s',
			'%s',
			NOW(),
			NOW()
		)
		RETURNING 
			id, 
			name, 
			email, 
			created_at, 
			updated_at;
	`
)
