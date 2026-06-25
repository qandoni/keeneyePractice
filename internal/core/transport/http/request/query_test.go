package core_http_request

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetIntQueryParam(t *testing.T) {
	req := httptest.NewRequest(
		"GET",
		"/students?limit=10",
		nil,
	)

	val, err := GetIntQueryParam(req, "limit")
	require.NoError(t, err)
	require.NotNil(t, val)
	require.Equal(t, 10, *val)
}

func TestGetIntQueryParam_InvalidValue(t *testing.T) {
	req := httptest.NewRequest(
		"GET",
		"/students?limit=abc",
		nil,
	)
	val, err := GetIntQueryParam(req, "limit")

	require.Error(t, err)
	require.Nil(t, val)
}

func TestGetIntQueryParam_Empty(t *testing.T) {
	req := httptest.NewRequest(
		"GET",
		"/students",
		nil,
	)

	val, err := GetIntQueryParam(req, "limit")

	require.NoError(t, err)
	require.Nil(t, val)
}

func TestGetIntPathValue(t *testing.T) {
	req := httptest.NewRequest(
		"GET",
		"/students/123",
		nil,
	)

	req.SetPathValue("id", "123")

	id, err := GetIntPathValue(req, "id")

	require.NoError(t, err)
	require.Equal(t, 123, id)
}

func TestGetIntPathValue_InvalidID(t *testing.T) {
	req := httptest.NewRequest(
		"GET",
		"/students/abc",
		nil,
	)

	req.SetPathValue("id", "abc")

	_, err := GetIntPathValue(req, "id")

	require.Error(t, err)
}
