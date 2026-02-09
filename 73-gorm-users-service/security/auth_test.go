package security_test

import (
	"testing"
	"user-service/security"

	"github.com/stretchr/testify/assert"
)

func TestGenerateAndValidateJWT(t *testing.T) {

	secret := "somesecret"
	username, email := "Jiten", "Jitenp@outlook.com"
	token, err := security.GenerateJWT(username, email, secret)

	assert.NoError(t, err)    // There should be no error
	assert.NotEmpty(t, token) // Some token to be returned

	t.Log(token) // This is to see what data has been generated

	claims, err := security.ValidateJWT(token, secret)
	assert.NoError(t, err) // There should be no error

	assert.NotNil(t, claims) // claims should not be nil

	// if claims == nil {
	// 	t.Fail()

	// }

	assert.Equal(t, "Jiten", claims.Username)
	assert.Equal(t, "Jitenp@outlook.com", claims.Email)
	assert.Equal(t, "user-service-app", claims.Issuer)

	assert.NotNil(t, claims.IssuedAt)
	assert.NotNil(t, claims.IssuedAt)

}
