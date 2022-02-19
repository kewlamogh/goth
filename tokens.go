package goth

import (
	"golang.org/x/crypto/bcrypt"
)

/* Creates a new Token from the given credentials. */
func GenToken(username string, password string) Token { 
	token, err := bcrypt.GenerateFromPassword([]byte(username+password), bcrypt.MinCost)
	checkError(err)
	tokenObject := Token{ Token: string(token) }

	return tokenObject
}

/* The struct representing the token. */
type Token struct {
	/* The token, which will be a SHA256 sum of the username and password concatenated. */
	Token string `json:"token"`
}

/* Checks if the token is valid. */
func (t Token) IsEqualToTokenOf(username string, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(t.Token), []byte(username+password))
	return err == nil, err
}