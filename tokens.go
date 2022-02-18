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
<<<<<<< HEAD
	err := bcrypt.CompareHashAndPassword([]byte(t.Token), []byte(username+password))	
=======
	err := bcrypt.CompareHashAndPassword([]byte(t.Token), []byte(username+password))
>>>>>>> 4ae93370717f6ed74443a072551e5c40b914c9ce
	return err == nil, err
}