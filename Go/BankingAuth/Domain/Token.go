package Domain

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
)

const HMAC_SAMPLE_SECRET = "7f6gh5df56g7bct8904fd7crl3fsj4892dfs3hdf"

type Claims struct {
	CustomerID string   `json:"customer_id"`
	Accounts   []string `json:"accounts"`
	Username   string   `json:"username"`
	Expiry     int64    `json:"expiry"`
	Role       string   `json:"role"`
}

func (c Claims) IsUserRole() bool {
	return c.Role == "user"
}

func BuildClaimsFromJwtMapClaims(mapClaims jwt.MapClaims) (*Claims, error) {
	bytes, err := json.Marshal(mapClaims)
	if err != nil {
		return nil, err
	}
	var c Claims
	err = json.Unmarshal(bytes, &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (c Claims) IsValidCustomerId(customerId string) bool {
	return c.CustomerID == customerId
}

func (c Claims) IsValidAccountId(accountId string) bool {
	if accountId != "" {
		accountFound := false
		for _, a := range c.Accounts {
			if a == accountId {
				accountFound = true
				break
			}
		}
		return accountFound
	}
	return true
}

func (c Claims) IsRequestVerifiedWithTokenClaims(urlParams map[string]string) bool {
	if c.CustomerID != urlParams["customer_id"] {
		return false
	}

	if !c.IsValidAccountId(urlParams["account_id"]) {
		return false
	}
	return true
}
