package kit

import (
	"testing"
)

func TestNodeMatchChild(t *testing.T) {
	root := &node{}
	root.children = []*node{
		{part: "a"},
	}
	if got := root.matchChild("a"); got == nil || got.part != "a" {
		t.Errorf("matchChild failed, want part 'a', got %v", got)
	}
	if got := root.matchChild("b"); got != nil {
		t.Errorf("matchChild failed, want nil, got %v", got)
	}

	root.children = append(root.children, &node{part: ":wild", isWild: true})
	if got := root.matchChild("b"); got == nil || !got.isWild {
		t.Errorf("matchChild failed, want wild, got %v", got)
	}
}

func TestNodeMatchChildren(t *testing.T) {
	root := &node{}
	root.children = []*node{
		{part: "a"},
		{part: ":wild", isWild: true},
	}
	got := root.matchChildren("a")
	if len(got) != 2 {
		t.Errorf("matchChildren failed, want 2, got %d", len(got))
	}
}

func TestNodeInsertAndSearch(t *testing.T) {
	root := &node{}
	root.insert("/p/:lang", []string{"p", ":lang"}, 0)
	root.insert("/p/go", []string{"p", "go"}, 0)

	tests := []struct {
		parts   []string
		wantPat string
	}{
		{[]string{"p", "go"}, "/p/go"},
		{[]string{"p", "java"}, "/p/:lang"},
		{[]string{"p"}, ""},
	}

	for _, tt := range tests {
		got := root.search(tt.parts, 0)
		if got == nil && tt.wantPat != "" {
			t.Errorf("search(%v) = nil, want %s", tt.parts, tt.wantPat)
		}
		if got != nil && got.pattern != tt.wantPat {
			t.Errorf("search(%v) = %s, want %s", tt.parts, got.pattern, tt.wantPat)
		}
	}
}
