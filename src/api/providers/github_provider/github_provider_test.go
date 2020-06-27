package github_provider

import (
	"io/ioutil"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAuthoriztionHeader(t *testing.T) {
	ioutil.NopCloser()
	header := getAuthorizationHeader("abc123")
	assert.EqualValues(t, "token abc123", header)
}
