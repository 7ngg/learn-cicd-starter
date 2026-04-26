package auth

import (
	"net/http"
	"testing"
)

func TestGetApiKeyCorrectHeader(t *testing.T) {
    tests := []struct {
        name     string
        header   http.Header
        expected string
        wantErr  bool
    }{
        {
            name:     "valid ApiKey header",
            header:   http.Header{"Authorization": []string{"ApiKey bootdev"}},
            expected: "bootdev",
            wantErr:  false,
        },
        {
            name:     "missing Authorization header",
            header:   http.Header{},
            expected: "",
            wantErr:  true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := GetAPIKey(tt.header)
            if (err != nil) != tt.wantErr {
                t.Errorf("unexpected error: %v", err)
            }
            if got != tt.expected {
                t.Errorf("got %q, want %q", got, tt.expected)
            }
        })
    }
}
