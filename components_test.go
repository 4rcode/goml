package goml_test

import (
	"encoding/xml"
	"testing"

	"github.com/4rcode/goml"
	"github.com/4rcode/gotu"
)

func TestDoctype(t *testing.T) {
	var assert = gotu.AssertWith(t)

	var (
		provided = goml.Doctype("HTML").String()
		expected = "<!DOCTYPE HTML>"
	)

	assert(provided == expected, provided, expected)
}

func TestElement(t *testing.T) {
	var assert = gotu.AssertWith(t)

	var provided, expected string

	provided = goml.Element("i", nil).String()
	expected = "<i></i>"

	assert(provided == expected, provided, expected)

	provided = goml.Element("i",
		goml.Attributes{
			{"class", "blue"}, {"style", "color: red;"}},
		goml.Text("text"),
	).String()
	expected = `<i class="blue" style="color: red;">text</i>`

	assert(provided == expected, provided, expected)
}

func TestFormat(t *testing.T) {
	var assert = gotu.AssertWith(t)

	var (
		provided = goml.Format("formatted %s", "string").String()
		expected = "formatted string"
	)

	assert(provided == expected, provided, expected)
}

func TestSequence(t *testing.T) {
	var assert = gotu.AssertWith(t)

	var (
		provided = goml.Sequence(
			goml.Text("SequenceOf"),
			goml.Text("Components"),
		).String()
		expected = "SequenceOfComponents"
	)

	assert(provided == expected, provided, expected)
}

func TestText(t *testing.T) {
	var assert = gotu.AssertWith(t)

	var (
		provided = goml.Text("text").String()
		expected = "text"
	)

	assert(provided == expected, provided, expected)
}

func TestToken(t *testing.T) {
	var assert = gotu.AssertWith(t)

	var (
		provided = goml.Token(xml.ProcInst{Target: "proc"}).String()
		expected = "<?proc?>"
	)

	assert(provided == expected, provided, expected)
}

func TestTokenWithError(t *testing.T) {
	var assert = gotu.AssertWith(t)

	var (
		provided = goml.Token("text").String()
		expected = ""
	)

	assert(provided == expected, provided, expected)
}

func TestWhen(t *testing.T) {
	var assert = gotu.AssertWith(t)

	var provided, expected string

	provided = goml.When(true, goml.Text("1")).String()
	expected = "1"

	assert(provided == expected, provided, expected)

	provided = goml.When(false, goml.Text("1")).String()
	expected = ""

	assert(provided == expected, provided, expected)

	provided = goml.When(true, goml.Text("1"), goml.Text("0")).String()
	expected = "1"

	assert(provided == expected, provided, expected)

	provided = goml.When(false, goml.Text("1"), goml.Text("0")).String()
	expected = "0"

	assert(provided == expected, provided, expected)
}
