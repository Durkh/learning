package Domain

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type AuthorizationRepo interface {
	FindBy(string, string) (*Login, error)
}

type AuthorizationDB struct {
	client *sqlx.DB
}

func (a AuthorizationDB) FindBy(username string, password string) (*Login, error) {

	var login Login

	sqlQuery := `SELECT username, u.customer_id, role, group_concat(a.account_id) as account_numbers FROM users u
                  LEFT JOIN accounts a ON a.customer_id = u.customer_id
                WHERE username = ? and password = ?
                GROUP BY a.customer_id`

	err := a.client.Get(&login, sqlQuery, username, password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("invalid credentials")
		} else {
			return nil, errors.New("unexpected database error")
		}
	}

	return &login, nil
}

func NewAuthorizationDB(client *sqlx.DB) AuthorizationDB {
	return AuthorizationDB{client: client}
}
