package doom

import (
	"io"

	"golang.org/x/net/html"
)

// Node is an alias for html.Node.
type Node html.Node

// A Filter accepts a Node and determines if it would match the criteria
// that the filter checks for.
type Filter func(node *Node) bool

// Parse is a wrapper around html.Parse that returns the resulting html.Node as
// a doom.Node instead. Errors returned by the underlying html.Parse call are
// propagated to the caller unmodified.
func Parse(r io.Reader) (*Node, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	return (*Node)(doc), nil
}

// Attribute indicates if the named attribute was set on the calling Node, and
// returns the value if applicable. In the case of boolean attributes, Attribute
// will return an empty string and true.
func (n *Node) Attribute(name string) (string, bool) {
	for _, attribute := range n.Attr {
		if attribute.Key == name {
			return attribute.Val, true
		}
	}

	return "", false
}

// Find will return the first child of the calling Node that matches all the
// provided filters. If no child node matches all the filters, it returns nil.
func (n *Node) Find(filters ...Filter) *Node {
	var walk func(node *html.Node) *html.Node

	walk = func(node *html.Node) *html.Node {
		for _, filter := range filters {
			if !filter((*Node)(node)) {
				goto fail
			}
		}

		return node

	fail:
		for next := node.FirstChild; next != nil; next = next.NextSibling {
			if result := walk(next); result != nil {
				return result
			}
		}

		return nil
	}

	return (*Node)(walk((*html.Node)(n)))
}

// Tag checks if a Node's element matches the provided tag.
func Tag(tag string) Filter {
	return func(node *Node) bool {
		return node.Type == html.ElementNode && node.Data == tag
	}
}

// AttributeEquals checks if a Node has the named attribute and if the value
// matches the provided value.
func AttributeEquals(attribute string, value string) Filter {
	return func(node *Node) bool {
		set, has := node.Attribute(attribute)
		return has && set == value
	}
}
