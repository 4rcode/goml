package goml

import (
	"encoding/xml"
	"net/http"
	"strings"
)

// [Attributes] is responsible for the definition of the [Element] attributes.
type Attributes [][2]string

// [Component] is responsible for rendering structured and reusable content,
// using a [Context].
//
// Component implements the [http.Handler] interface, therefore it can be
// directly used to render content over HTTP, using the standard [net/http]
// library.
//
// Component implements also the [fmt.Stringer] interface.
type Component func(Context)

// See [http.Handler]
func (c Component) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var ctx = Context{xml.NewEncoder(w)}

	if c != nil {
		c(ctx)
	}

	ctx.close()
}

// See [fmt.Stringer]
func (c Component) String() (s string) {
	var w strings.Builder
	var ctx = Context{xml.NewEncoder(&w)}

	if c != nil {
		c(ctx)
	}

	ctx.close()

	return w.String()
}
