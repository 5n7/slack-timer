package etc

import "testing"

func TestRemoveDuplicateSpace(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "basic",
			args: args{
				s: "foo  bar  baz",
			},
			want: "foo bar baz",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveDuplicateSpace(tt.args.s); got != tt.want {
				t.Errorf("RemoveDuplicateSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}
