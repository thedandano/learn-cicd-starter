package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name      string
		headers   http.Header
		want      string
		wantError bool
	}{
		{
			name:      "valid API key",
			headers:   http.Header{"Authorization": []string{"ApiKey test-key-123"}},
			want:      "test-key-12",
			wantError: false,
		},
		{
			name:      "missing authorization header",
			headers:   http.Header{},
			want:      "",
			wantError: true,
		},
		{
			name:      "malformed header - no space",
			headers:   http.Header{"Authorization": []string{"ApiKeytest-key"}},
			want:      "",
			wantError: true,
		},
		{
			name:      "wrong scheme",
			headers:   http.Header{"Authorization": []string{"Bearer test-key"}},
			want:      "",
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIKey(tt.headers)
			if (err != nil) != tt.wantError {
				t.Errorf("GetAPIKey() error = %v, wantError %v", err, tt.wantError)
				return
			}
			if got != tt.want {
				t.Errorf("GetAPIKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
