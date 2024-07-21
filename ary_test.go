package tern_test

import (
	"reflect"
	"testing"

	"github.com/karim-w/tern"
)

func TestAry(t *testing.T) {
	t.Parallel()
	type T = interface{}
	tests := []struct {
		name      string
		condition bool
		yes       T
		no        T
		want      T
	}{
		{
			name:      "true",
			condition: true,
			yes:       1,
			no:        2,
			want:      1,
		},
		{
			name:      "false",
			condition: false,
			yes:       1,
			no:        2,
			want:      2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tern.Ary(tt.condition, tt.yes, tt.no); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExample(
	t *testing.T,
) {
	res := tern.Ary(
		1 == 1,
		"yes",
		"no",
	)
	if res != "yes" {
		t.Errorf("got %v, want %v", res, "yes")
	}

	res = tern.Ary(
		1 == 2,
		"yes",
		"no",
	)
	if res != "no" {
		t.Errorf("got %v, want %v", res, "no")
	}
}
