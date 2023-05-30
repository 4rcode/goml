package goml_test

import (
	"encoding/xml"
	"log"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/4rcode/goml"
	"github.com/4rcode/gotu"
)

func TestComponentServeHTTP(t *testing.T) {
	var assert = gotu.AssertWith(t)

	var w = httptest.NewRecorder()

	goml.
		Text("text").
		ServeHTTP(w, nil)

	var (
		provided = w.Body.String()
		expected = "text"
	)

	assert(provided == expected, provided, expected)
}

func TestComponentServeHTTPWithError(t *testing.T) {
	var assert = gotu.AssertWith(t)

	var logs strings.Builder

	log.SetFlags(0)
	log.SetOutput(&logs)

	defer func() {
		log.SetOutput(os.Stderr)
	}()

	var w = httptest.NewRecorder()

	goml.Token(
		xml.StartElement{
			Name: xml.Name{
				Local: "element"}},
	).ServeHTTP(w, nil)

	var (
		provided = w.Body.String()
		expected = "<element>"
	)

	assert(provided == expected, provided, expected)

	provided = logs.String()
	expected = "unclosed tag <element>\n"

	assert(provided == expected, provided, expected)
}

func TestComponentString(t *testing.T) {
	var assert = gotu.AssertWith(t)

	var (
		provided = goml.Component(goml.Text("text")).String()
		expected = "text"
	)

	assert(provided == expected, provided, expected)
}

func TestComponentStringWithError(t *testing.T) {
	var assert = gotu.AssertWith(t)

	var logs strings.Builder

	log.SetFlags(0)
	log.SetOutput(&logs)

	defer func() {
		log.SetOutput(os.Stderr)
	}()

	var (
		provided = goml.Component(
			goml.Token(
				xml.StartElement{
					Name: xml.Name{
						Local: "element"}}),
		).String()
		expected = "<element>"
	)

	assert(provided == expected, provided, expected)

	provided = logs.String()
	expected = "unclosed tag <element>\n"

	assert(provided == expected, provided, expected)
}
