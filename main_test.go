package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGETPlayers(t *testing.T) {
	t.Run("returns Pepper's score", func(t *testing.T) {

		t.Run("test get date and time with valid endpoint", func(t *testing.T) {
			request, _ := http.NewRequest(http.MethodGet, "/datetime", nil)
			response := httptest.NewRecorder()
			GetDate(response, request)
			want := time.Now().Format(time.RFC3339)

			got := response.Body.String()

			assertSuccesMsg(t, response.Result().StatusCode)
			assertString(t, got, want)
		})
	})
}

func assertSuccesMsg(t *testing.T, got int) {
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
