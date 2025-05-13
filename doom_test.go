package doom_test

import (
	"strings"
	"testing"

	"github.com/jordanocokoljic/doom"
)

func TestNode_Text(t *testing.T) {
	body := `
<div>
	<h1 id="title">
		My
			Web
				Page
	</h1>
	<h2 id="subtitle">Welcome to my web page</h2>
</div>
	`

	root, err := doom.Parse(strings.NewReader(body))
	if err != nil {
		t.Fatalf("an error occured while parsing document: %s", err)
	}

	title := root.Find(doom.Tag("h1"))
	if title == nil {
		t.Fatal("failed to find node in tree with tag 'h1'")
	}

	text := title.Text()
	expected := "\n\t\tMy\n\t\t\tWeb\n\t\t\t\tPage\n\t"
	if text != expected {
		escapeNewlines := strings.Replace(text, "\n", "\\n", -1)
		escapeTabs := strings.Replace(escapeNewlines, "\t", "\\t", -1)
		t.Fatalf("incorrect text returned: %s", escapeTabs)
	}

	subtitle := root.Find(doom.Tag("h2"))
	if subtitle == nil {
		t.Fatalf("failed to find existent node with tag 'h2'")
	}

	text = subtitle.Text()
	expected = "Welcome to my web page"
	if text != expected {
		t.Fatalf("incorrect text returned: %s", text)
	}
}

func TestNode_Attribute(t *testing.T) {
	body := `
<div>
	<h1 id="title">My Web Page</h1>
	<h2 id="subtitle">Welcome to my web page</h2>
</div>
	`

	root, err := doom.Parse(strings.NewReader(body))
	if err != nil {
		t.Fatalf("an error occurred while parsing document: %s", err)
	}

	title := root.Find(doom.Tag("h1"))
	if title == nil {
		t.Fatal("failed to find node in tree with tag 'h1'")
	}

	id, ok := title.Attribute("id")
	if !ok {
		t.Fatal("node indicated it did not have 'id' attribute")
	}

	if id != "title" {
		t.Fatalf("node returned incorrect value for 'id' attribute: %s", id)
	}
}

func TestTag(t *testing.T) {
	body := `
<div>
	<h1 id="title">My Web Page</h1>
	<h2 id="subtitle">Welcome to my web page</h2>
</div>
	`

	root, err := doom.Parse(strings.NewReader(body))
	if err != nil {
		t.Fatalf("an error occured while parsing document: %s", err)
	}

	title := root.Find(doom.Tag("h1"))
	if title == nil {
		t.Fatal("failed to find node in tree with tag 'h1'")
	}

	fake := root.Find(doom.Tag("fake"))
	if fake != nil {
		t.Fatalf("incorrectly returned node with tag 'fake'")
	}
}

func TestAttributeEquals(t *testing.T) {
	body := `
<div>
	<h1 id="title">My Web Page</h1>
	<h2 id="subtitle">Welcome to my web page</h2>
</div>
	`

	root, err := doom.Parse(strings.NewReader(body))
	if err != nil {
		t.Fatalf("an error occured while parsing document: %s", err)
	}

	subtitle := root.Find(doom.AttributeEquals("id", "subtitle"))
	if subtitle == nil {
		t.Fatalf("failed to find node in tree with 'id=\"subtitle\"'")
	}

	fake := root.Find(doom.AttributeEquals("key", "value"))
	if fake != nil {
		t.Fatalf("incorrectly returned node for 'key=\"value\"'")
	}
}
