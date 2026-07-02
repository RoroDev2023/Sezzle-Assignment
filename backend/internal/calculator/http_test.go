package calculator

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCalculateEndpoint(t *testing.T) {
	t.Parallel()

	server := NewServer().Routes()

	tests := []struct {
		name       string
		body       string
		wantStatus int
		wantResult float64
	}{
		{
			name:       "valid calculation",
			body:       `{"operation":"multiply","a":6,"b":7}`,
			wantStatus: http.StatusOK,
			wantResult: 42,
		},
		{
			name:       "sqrt does not require b",
			body:       `{"operation":"sqrt","a":81}`,
			wantStatus: http.StatusOK,
			wantResult: 9,
		},
		{
			name:       "missing operand",
			body:       `{"operation":"add","a":1}`,
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "division by zero",
			body:       `{"operation":"divide","a":1,"b":0}`,
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "invalid operation",
			body:       `{"operation":"mod","a":1,"b":1}`,
			wantStatus: http.StatusNotFound,
		},
		{
			name:       "malformed json",
			body:       `{"operation":`,
			wantStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			request := httptest.NewRequest(http.MethodPost, "/api/calculate", bytes.NewBufferString(tt.body))
			request.Header.Set("Content-Type", "application/json")
			response := httptest.NewRecorder()

			server.ServeHTTP(response, request)

			if response.Code != tt.wantStatus {
				t.Fatalf("status = %d, want %d; body = %s", response.Code, tt.wantStatus, response.Body.String())
			}

			if tt.wantStatus != http.StatusOK {
				return
			}

			var payload calculateResponse
			if err := json.NewDecoder(response.Body).Decode(&payload); err != nil {
				t.Fatalf("decode response: %v", err)
			}

			if payload.Result != tt.wantResult {
				t.Fatalf("result = %v, want %v", payload.Result, tt.wantResult)
			}
		})
	}
}

func TestOperationsEndpoint(t *testing.T) {
	t.Parallel()

	request := httptest.NewRequest(http.MethodGet, "/api/operations", nil)
	response := httptest.NewRecorder()

	NewServer().Routes().ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", response.Code, http.StatusOK)
	}

	var payload map[string][]Operation
	if err := json.NewDecoder(response.Body).Decode(&payload); err != nil {
		t.Fatalf("decode response: %v", err)
	}

	if len(payload["operations"]) != len(Operations()) {
		t.Fatalf("operations = %v, want %v", payload["operations"], Operations())
	}
}

