package goth

import "testing"

func TestTokenMatcher(t *testing.T) {
	token := GenToken("bob", "dob")
	ok, err := token.IsEqualToTokenOf("job", "cob")

	if ok || err == nil {
		t.Errorf("should have thrown err, but didn't: %s", err.Error())
	}
}