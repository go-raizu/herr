package herr_test

import (
	"testing"

	"github.com/go-raizu/herr"
	"github.com/stretchr/testify/assert"
)

func TestStatusCode(t *testing.T) {
	type args struct {
		err herr.HTTPError
	}
	tt := []struct {
		name string
		args args
		want int
	}{
		{"", args{herr.ErrBadRequest}, 400},
		{"", args{herr.ErrNotFound}, 404},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equalf(t, tc.want, tc.args.err.StatusCode(), "StatusCode()")
		})
	}
}

func TestError(t *testing.T) {
	type args struct {
		err herr.HTTPError
	}
	tt := []struct {
		name string
		args args
		want string
	}{
		{"", args{herr.ErrBadRequest}, "400 (Bad Request)"},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equalf(t, tc.want, tc.args.err.Error(), "Error()")
		})
	}
}
