package healthz

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAlwaysUp(t *testing.T) {
	var (
		w = httptest.NewRecorder()
		r = httptest.NewRequest("GET", "/healthz", nil)
	)

	AlwaysUp(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("HTTP Status Code expected %d, got %d", http.StatusOK, w.Code)
	}

	if strings.Index(w.Body.String(), "background-color: green") <= 0 {
		t.Error("A green background is expected")
	}
}

func TestCheck(t *testing.T) {
	t.Run("Up", func(t *testing.T) {
		var (
			w = httptest.NewRecorder()
			r = httptest.NewRequest("GET", "/healthz", nil)
		)

		Check(func() bool { return true }).ServeHTTP(w, r)

		if w.Code != http.StatusOK {
			t.Errorf("HTTP Status Code expected %d, got %d", http.StatusOK, w.Code)
		}

		if strings.Index(w.Body.String(), "background-color: green") <= 0 {
			t.Error("A green background is expected")
		}
	})

	t.Run("Down", func(t *testing.T) {
		var (
			w = httptest.NewRecorder()
			r = httptest.NewRequest("GET", "/healthz", nil)
		)

		Check(func() bool { return false }).ServeHTTP(w, r)

		if w.Code != http.StatusServiceUnavailable {
			t.Errorf("HTTP Status Code expected %d, got %d", http.StatusServiceUnavailable, w.Code)
		}

		if strings.Index(w.Body.String(), "background-color: red") <= 0 {
			t.Error("A red background is expected")
		}
	})
}
