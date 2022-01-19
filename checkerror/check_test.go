package checkerror

import (
	"errors"
	"testing"
)

func TestCheckLog(t *testing.T) {
	type args struct {
		err  error
		tags []string
	}
	tests := []struct {
		name         string
		args         args
		wantHasError bool
	}{
		{
			name: "No Error",
			args: args{
				err:  nil,
				tags: []string{"no error"},
			},
			wantHasError: false,
		},
		{
			name: "Has Error",
			args: args{
				err:  errors.New("Unknown Error"),
				tags: []string{"TestError"},
			},
			wantHasError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotHasError := CheckLog(tt.args.err, tt.args.tags...); gotHasError != tt.wantHasError {
				t.Errorf("CheckLog() = %v, want %v", gotHasError, tt.wantHasError)
			}
		})
	}
}

func TestCheck(t *testing.T) {
	type args struct {
		err  error
		tags []string
	}
	tests := []struct {
		name string
		args args
	}{

		{
			name: "No Error",
			args: args{
				err:  nil,
				tags: []string{"no error"},
			},
		},
		{
			name: "Has Error",
			args: args{
				err:  errors.New("Unknown Error"),
				tags: []string{"TestError"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				err := recover()
				if tt.args.err != nil && err != nil {
					return
				}
				if err != nil {
					panic(err)
				}
			}()

			Check(tt.args.err, tt.args.tags...)
		})
	}
}
