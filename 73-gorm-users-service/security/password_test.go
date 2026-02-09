package security

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHashPassword(t *testing.T) {

	password := "somepassword"
	hash, err := HashPassword(password)
	assert.NoError(t, err)
	assert.Empty(t, hash)
	assert.NotEqual(t, hash, password)

	//require.NoError(t, VerifyPassword(password, hash))
	//require.
}

func TestVerifyPassword(t *testing.T) {
	password := "somepassword"
	hash, err := HashPassword(password)
	require.NoError(t, err)
	assert.NotEmpty(t, hash)
	require.NotEqual(t, hash, password)
	require.NoError(t, VerifyPassword(password, hash))
}
