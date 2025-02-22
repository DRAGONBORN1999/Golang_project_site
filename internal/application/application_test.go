package application_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DRAGONBORN1999/Golang_project_site/internal/application"
)

func sanitizeJSON(s string) string {
	return strings.Join(strings.Fields(s), "")
}

func TestCalcHandler(t *testing.T) {
	tests := []struct {
		name             string
		method           string
		body             interface{}
		expectedStatus   int
		expectedResponse string
	}{
		{
			name:             "Правильный запрос",
			method:           http.MethodPost,
			body:             map[string]string{"expression": "2 + 2"},
			expectedStatus:   http.StatusOK,
			expectedResponse: `{"result":"4"}`,
		},
		{
			name:             "Неправильный запрос",
			method:           http.MethodGet,
			body:             map[string]string{"name": "2 + 2"},
			expectedStatus:   http.StatusMethodNotAllowed,
			expectedResponse: `{"error":"Expression is not valid"}`,
		},
		{
			name:             "Некорректный запрос",
			method:           http.MethodPost,
			body:             "invalid body",
			expectedStatus:   http.StatusBadRequest,
			expectedResponse: `{"error":"Bad request"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var requestBody []byte
			if tt.body != nil {
				var err error
				requestBody, err = json.Marshal(tt.body)
				if err != nil {
					t.Fatal(err)
				}
			}

			reqPath := "/api/v1/calculate"

			req := httptest.NewRequest(tt.method, reqPath, bytes.NewBuffer(requestBody))
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(application.CalcHandler)

			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("Handler returned wrong status code: got %v want %v", status, tt.expectedStatus)
			}

			if tt.expectedResponse != "" {
				if sanitizedBody := sanitizeJSON(rr.Body.String()); sanitizedBody != sanitizeJSON(tt.expectedResponse) {
					t.Errorf("Handler returned unexpected response body: got '%v' want '%v'", rr.Body.String(), tt.expectedResponse)
				}
			}
		})
	}
}
