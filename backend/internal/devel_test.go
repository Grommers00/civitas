package internal_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/grommers00/civitas/backend/internal"
	"github.com/stretchr/testify/assert"
)

// ErrorTestHandler removes duplicate err handling logic
func ErrorTestHandler(err error, t *testing.T) {
	if err != nil {
		t.Errorf("Err: %d", err)
	}
}

func TestNotImplementedHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		internal.NotImplementedHandler("Test", w)
	}))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	ErrorTestHandler(err, t)

	results, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	ErrorTestHandler(err, t)

	if string(results) != "Test Not Implemented!\n" {
		t.Error("Expected result to be 'Test Not Implemented!\n'")
	}
}

func TestUnwrapJSONData(t *testing.T) {
	type TestJSONLoading struct {
		ID   int    `json:"id"`
		Game string `json:"game"`
	}

	t.Run("Valid PATH", func(t *testing.T) {
		ErrorTestHandler(internal.UnwrapJSONData("../mockdata/mocktests.json", &TestJSONLoading{}), t)
	})

	t.Run("Invalid PATH", func(t *testing.T) {
		err := internal.UnwrapJSONData("../badpath/mocktests.json", &TestJSONLoading{})
		assert.NotEqual(t, err, nil)
	})
}
