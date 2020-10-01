package routes_test

import (
	"net/http/httptest"
	"reflect"
	"testing"
)

func AssertExpected(t *testing.T, expected, received interface{}) {
	if !reflect.DeepEqual(expected, received) {
		t.Errorf("Got %v, wanted %v", received, expected)
	}
}

func SetupContext() *httptest.ResponseRecorder {
	w := httptest.NewRecorder()

	return w
}
