package github_provider

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAuthoriztionHeader(t *testing.T) {
	header := getAuthorizationHeader("abc123")
	assert.EqualValues(t, "token abc123", header)
}
