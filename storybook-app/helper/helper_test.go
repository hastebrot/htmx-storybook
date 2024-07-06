package helper

import (
	"strings"

	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"

	"testing"
)

func TestClasses(t *testing.T) {
	var tests = []struct {
		in  g.Node
		out string
	}{
		{Class("foo bar"), ` class="foo bar"`},
		{Class("[&foo] bar"), ` class="[&amp;foo] bar"`},
		{Classes(Class("foo"), Class("bar")), ` class="foo bar"`},
		{Classes(Class("[foo]"), Class("bar")), ` class="[foo] bar"`},
		{Classes(Class("[&foo]"), Class("bar")), ` class="[&amp;foo] bar"`},
	}

	for _, tt := range tests {
		if out := renderNode(tt.in); out != tt.out {
			t.Errorf("renderNode() is not '%s', actual '%s'", tt.out, out)
		}
	}
}

func renderNode(node g.Node) string {
	var builder strings.Builder
	_ = node.Render(&builder)
	return builder.String()
}
