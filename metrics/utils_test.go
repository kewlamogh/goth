package metrics

import (
	"testing"
)

func Test_checkError(t *testing.T) {
	type args struct {
		err error
	}

	tests := []struct {
		name     string
		args     args
		expected error
	}{
		{
			name: "nil error",
			args: args{
				err: nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checkError(tt.args.err)
		})
	}
}
