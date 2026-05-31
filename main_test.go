package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListItems(t *testing.T) {
	token, err := createToken("admin")
	if err != nil {
		t.Fatalf("create token: %v", err)
	}

	tests := []struct {
		name       string
		method     string
		token      string
		wantStatus int
		wantItems  int
	}{
		{
			name:       "missing token returns unauthorized",
			method:     http.MethodGet,
			wantStatus: http.StatusUnauthorized,
		},
		{
			name:       "invalid token returns unauthorized",
			method:     http.MethodGet,
			token:      "not-a-real-token",
			wantStatus: http.StatusUnauthorized,
		},
		{
			name:       "valid token returns items",
			method:     http.MethodGet,
			token:      token,
			wantStatus: http.StatusOK,
			wantItems:  3,
		},
		{
			name:       "wrong method returns method not allowed",
			method:     http.MethodPost,
			token:      token,
			wantStatus: http.StatusMethodNotAllowed,
		},
	}

	handler := appHandler()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, "/api/items", nil)
			if tt.token != "" {
				req.Header.Set("Authorization", "Bearer "+tt.token)
			}

			rec := httptest.NewRecorder()
			handler.ServeHTTP(rec, req)

			if rec.Code != tt.wantStatus {
				t.Fatalf("status = %d, want %d; body = %q", rec.Code, tt.wantStatus, rec.Body.String())
			}

			if tt.wantItems == 0 {
				return
			}

			var items []Item
			if err := json.NewDecoder(rec.Body).Decode(&items); err != nil {
				t.Fatalf("decode items: %v", err)
			}
			if len(items) != tt.wantItems {
				t.Fatalf("items length = %d, want %d", len(items), tt.wantItems)
			}
		})
	}
}
