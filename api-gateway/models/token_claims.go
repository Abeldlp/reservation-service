package models

type Claims struct {
	ID         int64  `json:"id"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	ValidToken bool   `json:"valid_token"`
}

func (c *Claims) Valid() error {
	c.ValidToken = true
	return nil
}
