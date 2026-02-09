package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserValidateSuccess(t *testing.T) {
	user := new(User)
	user.ID = 101
	user.Name = "JIten"
	user.Email = "Jitenp@outlook.com"
	user.Mobile = "9618558500"
	user.Password = "somepassword"

	err := user.Validate()
	// if err := user.Validate(); err != nil {
	// 	t.Fatalf("err should be nil %v", err.Error())
	// }
	assert.NoError(t, err)
}

func TestUserValidateFailure(t *testing.T) {
	user := new(User)
	user.ID = 101
	user.Email = "Jitenp@outlook.com"
	user.Mobile = "9618558500"
	user.Password = "somepassword"
	err := user.Validate()

	// if err := user.Validate(); err == nil {
	// 	t.Fatalf("err should not be  nil %v", err.Error())
	// }

	assert.Error(t, err)
}

func TestUserSame(t *testing.T) {
	user1 := &User{}
	user1.ID = 101
	user1.Email = "Jitenp@outlook.com"
	user1.Mobile = "9618558500"
	user1.Password = "somepassword"

	// user2 := User{}
	// user2.ID = 101
	// user2.Email = "Jitenp@outlook.com"
	// user2.Mobile = "9618558500"
	// user2.Password = "somepassword"
	ref := user1

	assert.Same(t, user1, ref)

	slice1 := []int{10, 11, 12}
	slice2 := &slice1

	assert.Same(t, &slice1, slice2)

}

func TestUserEqual(t *testing.T) {
	user1 := User{}
	user1.ID = 101
	user1.Email = "Jitenp@outlook.com"
	user1.Mobile = "9618558500"
	user1.Password = "somepassword"

	user2 := User{}
	user2.ID = 101
	user2.Email = "Jitenp@outlook.com"
	user2.Mobile = "9618558500"
	user2.Password = "somepassword"

	assert.Equal(t, user1, user2)

	slice1 := []int{10, 11, 12}

	slice2 := []int{10, 11, 12}

	assert.Equal(t, slice1, slice2)

}

// go test -test.fullpath=true -timeout 30s user-service/models
// go test user-service/models -v
