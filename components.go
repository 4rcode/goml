package goml

import (
	"encoding/xml"
	"fmt"
)

func Doctype(doctype string) Component {
	return Token(xml.Directive("DOCTYPE " + doctype))
}

func Element(
	name string,
	attributes [][2]string,
	children ...func(Context),
) Component {
	var attrs []xml.Attr

	if l := len(attributes); l > 0 {
		attrs = make([]xml.Attr, 0, l)

		for _, attribute := range attributes {
			attrs = append(attrs, xml.Attr{
				Name:  xml.Name{Local: attribute[0]},
				Value: attribute[1]})
		}
	}

	return Sequence(
		Token(xml.StartElement{
			Name: xml.Name{Local: name},
			Attr: attrs}),
		Sequence(children...),
		Token(xml.EndElement{
			Name: xml.Name{Local: name}}))
}

func Format(format string, arguments ...interface{}) Component {
	return Text(
		fmt.Sprintf(format, arguments...))
}

func Sequence(components ...func(Context)) Component {
	return func(ctx Context) {
		for _, component := range components {
			if component != nil {
				component(ctx)
			}
		}
	}
}

func Text(text string) Component {
	return Token(xml.CharData(text))
}

func Token(token interface{}) Component {
	return func(ctx Context) {
		ctx.encode(token)
	}
}

func When(
	isTrue bool,
	then func(Context),
	otherwise ...func(Context),
) Component {
	if isTrue {
		return then
	}

	return Sequence(otherwise...)
}
