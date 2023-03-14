package swift

import "testing"

func TestCreateIdentifierName(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"No translation",
			args{"Foo"},
			"Foo",
		},
		{
			"Remove spaces",
			args{"Foo bar"},
			"FooBar",
		},
		{
			"Replace punctuation with underscores",
			args{"Foo!bar"},
			"Foo_bar",
		},
		{
			"Replace bracing with underscores (smartly)",
			args{"Foo(bar)"},
			"Foo_bar",
		},
		{
			"Empty string",
			args{""},
			"",
		},
		{
			"Underscore string",
			args{"_"},
			"_",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateIdentifierName(tt.args.input); got != tt.want {
				t.Errorf("CreateIdentifierName() = %v, want %v", got, tt.want)
			}
		})
	}
}
