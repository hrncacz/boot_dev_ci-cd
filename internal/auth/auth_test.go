package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetAuth(t *testing.T) {
	tests := []struct {
		name  string
		input map[string]string
		want  string
	}{
		{
			name: "valid input",
			input: map[string]string{
				"Authorization": "ApiKey FOIDSFsdjfs09434hjkšč4ščšč",
			},
			want: "FOIDSFsdjfs09434hjkšč4ščšč",
		},
		{
			name: "AppKey instead of ApiKey input",
			input: map[string]string{
				"Authorization": "AppKey FOIDSFsdjfs09434hjkšč4ščšč",
			},
			want: "",
		},
		{
			name: "no ApiKey prefix",
			input: map[string]string{
				"Authorization": "FOIDSFsdjfs09434hjkšč4ščšč",
			},
			want: "",
		},
		{
			name:  "missing Authentication header",
			input: map[string]string{},
			want:  "",
		},
	}

	for _, tc := range tests {
		h := http.Header{}
		for key, val := range tc.input {
			h.Set(key, val)
		}
		got, _ := GetAPIKey(h)

		if !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
		}
	}
}
