package domain

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetUserNoUserFound(t *testing.T) {
	user, err := UserDao.GetUser(0)
	assert.Nil(t, user, "we were not expecting a user with id 0")
	// if user != nil {
	// 	t.Error("we were not expecting a user with id 0")
	// }

	assert.NotNil(t, err, "we were not expecting a user with id 0")
	// if err == nil {
	// 	t.Error("we were expecting an error when user id is 0")
	// }

	assert.EqualValues(t, http.StatusNotFound, err.StatusCode, "we were not expecting a user with id 0")
	// if err.StatusCode != http.StatusNotFound {
	// 	t.Error("we were expecting 404 when user is not found")
	// }
	assert.EqualValues(t, "not_found", err.Code)

}

func TestGetUserNoError(t *testing.T) {
	user, err := UserDao.GetUser(123)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 123, user.ID)
	assert.EqualValues(t, "Stephen", user.FirstName)
	assert.EqualValues(t, "Curry", user.LastName)
	assert.EqualValues(t, "curry@gmail.com", user.Email)
}
