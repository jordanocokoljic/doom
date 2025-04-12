package doom_test

import (
	"strings"
	"testing"

	"github.com/jordanocokoljic/doom"
)

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
		t.Fatalf("failed to find existent node with tag 'h1'")
	}

	fake := root.Find(doom.Tag("fake"))
	if fake != nil {
		t.Fatalf("found non-existent node with tag 'fake'")
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
		t.Fatalf("failed to find existent node with attribute 'id=\"subtitle\"'")
	}

	fake := root.Find(doom.AttributeEquals("key", "value"))
	if fake != nil {
		t.Fatalf("found non-existent node with attribute 'key=\"value\"'")
	}
}
