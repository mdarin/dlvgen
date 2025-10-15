package finder_test

import (
	"dlvgen/internal/finder"
	"testing"
)

func TestFindMainProgram(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		paths []string
		want  string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := finder.FindMainProgram(tt.paths)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("FindMainProgram() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectBestCandidate(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		hits []string
		want string
	}{
		{
			name: "[cmd/app/main.go,main.go,pkg/go-find/cmd/go-find/main.go]",
			hits: []string{"cmd/app/main.go", "main.go", "pkg/go-find/cmd/go-find/main.go"},
			want: "main.go",
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := finder.SelectBestCandidate(tt.hits)
			if tt.want != got {
				t.Errorf("SelectBestCandidate() = %v, want %v", got, tt.want)
			}
		})
	}
}
