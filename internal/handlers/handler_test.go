package handlers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	testCases := []struct {
		name         string
		url          string
		expectedCode int
		expectedBody string
	}{
		{
			name:         "Success",
			url:          "/greet?name=Alice",
			expectedCode: http.StatusOK,
			expectedBody: "Hello Alice",
		},
	}

	h := NewHandler()

	for _, tc := range testCases {
		req := httptest.NewRequest(http.MethodGet, tc.url, nil)

		rr := httptest.NewRecorder()

		h.Greet(rr, req)

		if rr.Code != tc.expectedCode {
			t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, tc.expectedCode)
		}

		body, err := io.ReadAll(rr.Body)
		if err != nil {
			t.Fatalf("could not read response body: %v", err)
		}

		if string(body) != tc.expectedBody {
			t.Errorf("handler returned unexpected body: got %q wannt %q", string(body), tc.expectedBody)
		}
	}
}
