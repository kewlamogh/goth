package goth

import (
	"crypto/sha256"
	"errors"
	"fmt"
)

/* Creates a new Token from the given credentials. */
func GenToken(username string, password string) Token { 
	token := sha256.Sum256([]byte(username+password))
	tokenObject := Token{ Token: fmt.Sprintf("%x", token) }

	return tokenObject
}

/* The struct representing the token. */
type Token struct {
	/* The token, which will be a SHA256 sum of the username and password concatenated. */
	Token string `json:"token"`
}

/* Checks if the token is valid. */
func (t Token) IsEqualToTokenOf(username string, password string) (bool, error) {
	expectedToken := GenToken(username, password)
	if expectedToken.Token != t.Token {
		return false, errors.New("wrong")
	}

	return true, nil
}
