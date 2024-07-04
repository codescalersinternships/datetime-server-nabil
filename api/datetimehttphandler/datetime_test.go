package datetime

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGETDate(t *testing.T) {
	t.Run("test valid date case", func(t *testing.T) {

		request, _ := http.NewRequest(http.MethodGet, "/datetime", nil)
		response := httptest.NewRecorder()
		GetDate(response, request)
		want := time.Now().Format("2003-29-05 01:02:00")

		got := response.Body.String()

		assertSuccesMsg(t, response.Result().StatusCode)
		assertString(t, got, want)
	})
	t.Run("wrong url", func(t *testing.T) {

		request, _ := http.NewRequest(http.MethodGet, "/invalid", nil)
		response := httptest.NewRecorder()
		GetDate(response, request)

		got := response.Result().StatusCode

		assertNotFoundMsg(t, got)
	})
	t.Run("wrong method", func(t *testing.T) {

		request, _ := http.NewRequest(http.MethodPost, "/datetime", nil)
		response := httptest.NewRecorder()
		GetDate(response, request)

		got := response.Result().StatusCode

		assertUnallowedMethod(t, got)
	})
}

func assertUnallowedMethod(t *testing.T, got int) {
	t.Helper()
	if got != http.StatusMethodNotAllowed {
		t.Errorf("expected status code 405 but got %d", got)
	}
}
func assertNotFoundMsg(t *testing.T, got int) {
	t.Helper()
	if got != http.StatusNotFound {
		t.Errorf("expected status code 404 but got %d", got)
	}
}

func assertSuccesMsg(t *testing.T, got int) {
	t.Helper()
	if got != 200 {
		t.Errorf("expected status code 200 but got %d", got)
	}
}
func assertString(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
