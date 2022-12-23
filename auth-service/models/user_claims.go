package models

type Claims struct {
	Username   string `json:"username"`
	Email      string `json:"email"`
	ValidToken bool   `json:"valid_token"`
}

func (c *Claims) Valid() error {
	c.ValidToken = true
	return nil
}
