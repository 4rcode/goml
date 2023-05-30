package goml

import (
	"encoding/xml"
	"log"
)

// [Context] is the internal context for rendering [Component] instances.
type Context struct {
	encoder *xml.Encoder
}

func (c Context) close() {
	if err := c.encoder.Close(); err != nil {
		log.Println(err)
	}
}

func (c Context) encode(token interface{}) {
	if err := c.encoder.EncodeToken(token); err != nil {
		log.Println(err)
	}
}
