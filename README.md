# Doom
`doom` is a DOM utility library for Go that builds on-top of the `/x/net/html`
library providing helper functions for finding specific elements within the
tree.

## Example Usage

```go
package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/jordanocokoljic/doom"
)

func main() {
	body := `
<div>
    <h1 id="title">My Web Page</h1>
    <h2 id="subtitle">Welcome to my web page</h2>
</div>
	`

	doc, err := doom.Parse(strings.NewReader(body))
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to parse provided HTML")
		os.Exit(1)
	}

	h1 := doc.Find(doom.AttributeEquals("id", "title"))
	// h1 is a Node parsed from <h1 id="title">My Web Page</h1>
}
```
