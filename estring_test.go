package estring

import "testing"

func TestUcFirst(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"t1", args{"ha"}, "Ha"},
		{"t2", args{"Ha"}, "Ha"},
		{"t3", args{"哈Ha"}, "哈Ha"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UcFirst(tt.args.s); got != tt.want {
				t.Errorf("UcFirst() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLcFirst(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"t1", args{"ha"}, "ha"},
		{"t2", args{"Ha"}, "ha"},
		{"t3", args{"哈Ha"}, "哈Ha"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LcFirst(tt.args.s); got != tt.want {
				t.Errorf("LcFirst() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStudly(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"t1", args{"this_is_a_demo"}, "ThisIsADemo"},
		{"t2", args{"this.is.a.demo"}, "ThisIsADemo"},
		{"t3", args{"this-is-a-demo"}, "ThisIsADemo"},
		{"t4", args{"this.is-a_demo"}, "ThisIsADemo"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Studly(tt.args.s); got != tt.want {
				t.Errorf("Studly() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCamel(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"t1", args{"this_is_a_demo"}, "thisIsADemo"},
		{"t2", args{"this.is.a.demo"}, "thisIsADemo"},
		{"t3", args{"this-is-a-demo"}, "thisIsADemo"},
		{"t4", args{"this.is-a_demo"}, "thisIsADemo"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Camel(tt.args.s); got != tt.want {
				t.Errorf("Camel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPadLeft(t *testing.T) {
	type args struct {
		s      string
		length int
		pad    string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"t1", args{"a", 5, "B"}, "BBBBa"},
		{"t2", args{"a", 0, "B"}, "a"},
		{"t3", args{"aaaaa", 5, "B"}, "aaaaa"},
		{"t4", args{"aaaaa", 4, "B"}, "aaaaa"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PadLeft(tt.args.s, tt.args.length, tt.args.pad); got != tt.want {
				t.Errorf("PadLeft() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPadRight(t *testing.T) {
	type args struct {
		s      string
		length int
		pad    string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"t1", args{"a", 5, "B"}, "aBBBB"},
		{"t2", args{"a", 0, "B"}, "a"},
		{"t3", args{"aaaaa", 5, "B"}, "aaaaa"},
		{"t4", args{"aaaaa", 4, "B"}, "aaaaa"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PadRight(tt.args.s, tt.args.length, tt.args.pad); got != tt.want {
				t.Errorf("PadRight() = %v, want %v", got, tt.want)
			}
		})
	}
}
