package x_test

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/4rcode/goml"
	x "github.com/4rcode/goml/symbols"
)

func Example() {
	var r *http.Request
	var w = httptest.NewRecorder()

	// 1. Import the default symbols

	/*

		import x "github.com/4rcode/goml/symbols"

	*/

	// 2. Build your own components

	var CustomComponent = func(labels ...string) goml.Component {
		return func(ctx goml.Context) {
			for _, label := range labels {
				x.W(label != "d",
					x.T(label))(ctx)
			}
		}
	}

	var Component = func(text string) goml.Component {
		return x.S(
			x.D("html"),
			x.E("html", nil,
				x.T("\n"),
				x.E("head", nil,
					x.E("link",
						x.A{{"rel", "stylesheet"}, {"href", "style.css"}})),
				x.T("\n"),
				x.E("body", nil,
					goml.Token(
						xml.Comment("comment")),
					x.W(1 < 2,
						x.T("true"),
						x.T("or "),
						x.T("false?")),
					x.E("h1", nil,
						x.T("title")),
					x.E("p", nil,
						x.F("short %s", text)),
					x.T("\n"),
					func(ctx goml.Context) {
						x.E("div", nil,
							x.T("inline component"))(ctx)
					},
					x.T("\n"),
					x.E("div", nil,
						x.T("custom component "),
						CustomComponent("a", "b", "c", "d")),
					x.T("\n"))))
	}

	// 3. Initialise and serve the component

	var component = Component("text")

	component.ServeHTTP(w, r)

	// 4. Check the response body

	fmt.Print(w.Body)
	// output:
	// <!DOCTYPE html><html>
	// <head><link rel="stylesheet" href="style.css"></link></head>
	// <body><!--comment-->true<h1>title</h1><p>short text</p>
	// <div>inline component</div>
	// <div>custom component abc</div>
	// </body></html>
}
