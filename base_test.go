package herr

import (
	"errors"
	"testing"
)

func TestHTTPError_Is(t *testing.T) {
	type args struct {
		a, b HTTPError
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{ErrBadRequest, ErrBadRequest}, true},
		{"", args{ErrBadRequest, ErrNotFound}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errors.Is(tt.args.a, tt.args.b)
		})
	}
}
