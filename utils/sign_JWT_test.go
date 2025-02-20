package utils

import (
	"testing"
)

func TestSignJWT(t *testing.T) {
	userMail := "test@example.com"
	userPass := "testpassword"
	token, err := SignJWT(&userMail, &userPass)
	if err != nil {
		t.Error(err)
	}
	t.Log(*token)
}
